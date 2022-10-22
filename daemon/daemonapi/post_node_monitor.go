package daemonapi

import (
	"net/http"

	"github.com/goccy/go-json"

	"opensvc.com/opensvc/core/cluster"
	"opensvc.com/opensvc/daemon/msgbus"
	"opensvc.com/opensvc/util/hostname"
	"opensvc.com/opensvc/util/pubsub"
)

func (a *DaemonApi) PostNodeMonitor(w http.ResponseWriter, r *http.Request) {
	var (
		payload      PostNodeMonitor
		validRequest bool
	)
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	cmd := msgbus.SetNmon{
		Node:    hostname.Hostname(),
		Monitor: cluster.NodeMonitor{},
	}
	if payload.GlobalExpect != nil {
		validRequest = true
		cmd.Monitor.GlobalExpect = *payload.GlobalExpect
	}
	if payload.State != nil {
		validRequest = true
		cmd.Monitor.Status = *payload.State
	}
	if !validRequest {
		sendError(w, http.StatusBadRequest, "need at least state or global_expect")
		return
	}
	bus := pubsub.BusFromContext(r.Context())
	msgbus.PubSetNmon(bus, cmd)
	response := ResponseInfoStatus{
		Info:   0,
		Status: "instance monitor pushed pending ops",
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}