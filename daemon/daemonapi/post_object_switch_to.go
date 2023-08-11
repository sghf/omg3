package daemonapi

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/opensvc/om3/core/instance"
	"github.com/opensvc/om3/core/path"
	"github.com/opensvc/om3/daemon/api"
	"github.com/opensvc/om3/daemon/msgbus"
	"github.com/opensvc/om3/util/hostname"
	"github.com/opensvc/om3/util/pubsub"
)

func (a *DaemonApi) PostObjectSwitchTo(ctx echo.Context) error {
	var (
		payload = api.PostObjectSwitchTo{}
		value   = instance.MonitorUpdate{}
		p       path.T
		err     error
	)
	if err := ctx.Bind(&payload); err != nil {
		return JSONProblem(ctx, http.StatusBadRequest, "Invalid Body", err.Error())
	}
	p, err = path.Parse(payload.Path)
	if err != nil {
		return JSONProblemf(ctx, http.StatusBadRequest, "Invalid Body", "Invalid path: %s", payload.Path)
	}
	globalExpect := instance.MonitorGlobalExpectPlacedAt
	options := instance.MonitorGlobalExpectOptionsPlacedAt{}
	options.Destination = append(options.Destination, payload.Destination...)
	value = instance.MonitorUpdate{
		GlobalExpect:             &globalExpect,
		GlobalExpectOptions:      options,
		CandidateOrchestrationId: uuid.New(),
	}
	msg := msgbus.SetInstanceMonitor{
		Path:  p,
		Node:  hostname.Hostname(),
		Value: value,
		Err:   make(chan error),
	}
	a.EventBus.Pub(&msg, pubsub.Label{"path", p.String()}, labelApi)
	var errs error
	for {
		select {
		case err := <-msg.Err:
			if err != nil {
				errs = errors.Join(errs, err)
			} else if errs != nil {
				return JSONProblemf(ctx, http.StatusBadRequest, "set monitor", "%s", errs)
			} else {
				return ctx.JSON(http.StatusOK, api.MonitorUpdateQueued{
					OrchestrationId: value.CandidateOrchestrationId,
				})
			}
		case <-ctx.Request().Context().Done():
			return JSONProblemf(ctx, http.StatusGone, "set monitor", "")
		}
	}
}
