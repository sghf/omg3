package daemondata

import (
	"context"
	"reflect"
	"runtime"
	"time"

	"github.com/rs/zerolog"

	"opensvc.com/opensvc/core/cluster"
	"opensvc.com/opensvc/core/hbtype"
	"opensvc.com/opensvc/daemon/daemonctx"
	"opensvc.com/opensvc/daemon/daemonlogctx"
	"opensvc.com/opensvc/util/callcount"
	"opensvc.com/opensvc/util/durationlog"
	"opensvc.com/opensvc/util/jsondelta"
	"opensvc.com/opensvc/util/pubsub"
)

type (
	caller interface {
		call(context.Context, *data)
	}

	data struct {
		// previousRemoteInfo map[node] of remoteInfo from pending data just
		// after commit, it is used to publish diff for other nodes
		previousRemoteInfo map[string]remoteInfo

		// pending is the live current data (after apply patch, commit local pendingOps)
		pending *cluster.Status

		pendingOps []jsondelta.Operation // local data pending operations not yet in patchQueue
		patchQueue patchQueue            // local data patch queue for remotes
		gen        uint64                // gen of local TNodeData
		hbMsgType  string                // latest created hb message type
		localNode  string
		counterCmd chan<- interface{}
		log        zerolog.Logger
		bus        *pubsub.Bus

		// msgLocalGen hold the latest published msg gen for localhost
		msgLocalGen map[string]uint64
		hbSendQ     chan<- hbtype.Msg
	}

	gens       map[string]uint64
	patchQueue map[string]jsondelta.Patch

	// remoteInfo struct holds information about remote node used to publish diff
	remoteInfo struct {
		nmonUpdated       time.Time
		nodeStatus        cluster.NodeStatus
		smonUpdated       map[string]time.Time
		instCfgUpdated    map[string]time.Time
		instStatusUpdated map[string]time.Time
		gen               uint64
	}
)

var (
	cmdDurationWarn = time.Second

	// propagationInterval is the minimum interval of:
	// - commit pending ops (update patch queue, send local events to event.Event subscribers)
	// - pub applied changes from peers
	// - queueNewHbMsg (hb message type change, push msg to hb send queue)
	propagationInterval = 250 * time.Millisecond

	// subHbRefreshInterval is the minimum interval for update of: sub.hb
	subHbRefreshInterval = 10 * time.Second

	countRoutineInterval = 1 * time.Second
)

func run(ctx context.Context, cmdC <-chan interface{}) {
	counterCmd, cancel := callcount.Start(ctx, idToName)
	defer cancel()
	d := newData(counterCmd)
	d.log = daemonlogctx.Logger(ctx).With().Str("name", "daemondata").Logger()
	d.log.Info().Msg("starting")
	defer d.log.Info().Msg("stopped")
	d.bus = pubsub.BusFromContext(ctx)

	watchCmd := &durationlog.T{Log: d.log}
	watchDurationCtx, watchDurationCancel := context.WithCancel(context.Background())
	defer watchDurationCancel()
	var beginCmd = make(chan interface{})
	var endCmd = make(chan bool)
	go func() {
		watchCmd.WarnExceeded(watchDurationCtx, beginCmd, endCmd, cmdDurationWarn, "data")
	}()

	propagationTicker := time.NewTicker(propagationInterval)
	defer propagationTicker.Stop()

	subHbRefreshTicker := time.NewTicker(subHbRefreshInterval)
	defer subHbRefreshTicker.Stop()
	d.msgLocalGen = make(map[string]uint64)

	countRoutineTicker := time.NewTicker(countRoutineInterval)
	defer countRoutineTicker.Stop()

	d.hbSendQ = daemonctx.HBSendQ(ctx)

	for {
		select {
		case <-ctx.Done():
			bg, cleanupCancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			go func() {
				d.log.Debug().Msg("drop pending cmds")
				defer cleanupCancel()
				for {
					select {
					case c := <-cmdC:
						dropCmd(ctx, c)
					case <-bg.Done():
						d.log.Debug().Msg("drop pending cmds done")
						return
					}
				}
			}()

			return
		case <-propagationTicker.C:
			changes := d.commitPendingOps()
			if !changes && !gensEqual(d.msgLocalGen, d.pending.Cluster.Node[d.localNode].Status.Gen) {
				changes = true
			}
			d.pubPeerDataChanges()
			select {
			case <-subHbRefreshTicker.C:
				d.setSubHb()
				changes = true
			case <-countRoutineTicker.C:
				d.pending.Monitor.Routines = runtime.NumGoroutine()
			default:
			}
			if changes {
				if err := d.queueNewHbMsg(); err != nil {
					d.log.Error().Err(err).Msg("queue hb message")
				}
			}
			propagationTicker.Reset(propagationInterval)
		case cmd := <-cmdC:
			if c, ok := cmd.(caller); ok {
				beginCmd <- cmd
				c.call(ctx, d)
				endCmd <- true
			} else {
				d.log.Debug().Msgf("%s{...} is not a caller-interface cmd", reflect.TypeOf(cmd))
				counterCmd <- idUndef
			}
		}
	}
}

type (
	errorSetter interface {
		setError(context.Context, error)
	}

	doneSetter interface {
		setDone(context.Context, bool)
	}
)

// dropCmd drops commands with side effects
func dropCmd(ctx context.Context, c interface{}) {
	// TODO implement all side effects
	switch cmd := c.(type) {
	case errorSetter:
		cmd.setError(ctx, nil)
	case doneSetter:
		cmd.setDone(ctx, true)
	}
}

func gensEqual(a, b gens) bool {
	if len(a) != len(b) {
		return false
	} else {
		for n, v := range a {
			if b[n] != v {
				return false
			}
		}
	}
	return true
}
