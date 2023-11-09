package daemonapi

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/opensvc/om3/util/command"
)

func (a *DaemonApi) PostDaemonRestart(ctx echo.Context) error {
	log := LogHandler(ctx, "PostDaemonRestart")
	log.Infof("starting")

	execname, err := os.Executable()
	if err != nil {
		return JSONProblemf(ctx, http.StatusInternalServerError, "Server error", "can't detect om execname: %s", err)
	}

	cmd := command.New(
		command.WithName(execname),
		command.WithArgs([]string{"daemon", "restart"}),
	)

	err = cmd.Start()
	if err != nil {
		log.Errorf("called StartProcess: %s", err)
		return JSONProblemf(ctx, http.StatusInternalServerError, "Server error", "daemon restart failed: %s", err)
	}
	log.Infof("called daemon restart")
	return JSONProblem(ctx, http.StatusOK, "background daemon restart has been called", "")
}