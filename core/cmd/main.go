package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/core/event"
	"github.com/opensvc/om3/core/naming"
	"github.com/opensvc/om3/daemon/msgbus"
)

func WaitInstanceMonitor(ctx context.Context, c *client.T, p naming.Path, timeout time.Duration, errC chan error) error {
	var (
		err      error
		evReader event.ReadCloser
	)
	filters := []string{"InstanceMonitorUpdated,path=" + p.String()}
	getEvents := c.NewGetEvents().SetFilters(filters).SetLimit(1)
	if timeout > 0 {
		getEvents = getEvents.SetDuration(timeout)
	}
	evReader, err = getEvents.GetReader()
	if err != nil {
		return err
	}

	if x, ok := evReader.(event.ContextSetter); ok {
		x.SetContext(ctx)
	}
	go func() {
		defer func() {
			if err != nil {
				err = fmt.Errorf("wait instance monitor update failed on object %s: %w", p, err)
			}
			select {
			case <-ctx.Done():
			case errC <- err:
			}
		}()

		go func() {
			// close reader when ctx is done
			select {
			case <-ctx.Done():
				_ = evReader.Close()
			}
		}()
		for {
			ev, readError := evReader.Read()
			if readError != nil {
				if errors.Is(readError, io.EOF) {
					err = fmt.Errorf("no more events, wait %v failed %s: %w", p, time.Now(), err)
				} else {
					err = readError
				}
				return
			}
			_, err = msgbus.EventToMessage(*ev)
			return
		}
	}()
	return nil
}