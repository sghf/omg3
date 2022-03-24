package hbucast

import (
	"context"
	"encoding/json"
	"net"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"opensvc.com/opensvc/core/hbtype"
	"opensvc.com/opensvc/daemon/daemonctx"
	"opensvc.com/opensvc/daemon/hb/hbctrl"
	"opensvc.com/opensvc/daemon/listener/encryptconn"
)

type (
	// rx holds an hb unicast receiver
	rx struct {
		ctx     context.Context
		id      string
		nodes   []string
		addr    string
		port    string
		intf    string
		timeout time.Duration

		name   string
		log    zerolog.Logger
		cmdC   chan<- interface{}
		msgC   chan<- *hbtype.Msg
		cancel func()
	}
)

// Id implements the Id function of the Receiver interface for rx
func (r *rx) Id() string {
	return r.id
}

// Stop implements the Stop function of the Receiver interface for rx
func (r *rx) Stop() error {
	r.cancel()
	for _, node := range r.nodes {
		r.cmdC <- hbctrl.CmdDelWatcher{
			HbId:     r.id,
			Nodename: node,
		}
	}
	return nil
}

// Start implements the Start function of the Receiver interface for rx
func (r *rx) Start(cmdC chan<- interface{}, msgC chan<- *hbtype.Msg) error {
	ctx, cancel := context.WithCancel(r.ctx)
	r.cmdC = cmdC
	r.msgC = msgC
	r.cancel = cancel
	r.log.Info().Msg("starting")
	listener, err := net.Listen("tcp", r.addr+":"+r.port)
	if err != nil {
		r.log.Error().Err(err).Msg("listen failed")
		return err
	}
	r.log.Info().Msgf("listen on %s", r.addr+":"+r.port)
	started := make(chan bool)
	go func() {
		for _, node := range r.nodes {
			cmdC <- hbctrl.CmdAddWatcher{
				HbId:     r.id,
				Nodename: node,
				Ctx:      ctx,
				Timeout:  r.timeout,
			}
		}
		go func() {
			select {
			case <-ctx.Done():
				r.log.Info().Msg("closing " + r.addr)
				_ = listener.Close()
				r.log.Info().Msg("closed " + r.addr)
				r.cancel()
				return
			}
		}()
		started <- true
		for {
			conn, err := listener.Accept()
			if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					break
				} else {
					r.log.Error().Err(err).Msg("Accept")
					continue
				}
			}
			if err := conn.SetDeadline(time.Now().Add(r.timeout)); err != nil {
				r.log.Info().Err(err).Msg("SetReadDeadline")
				continue
			}
			clearConn := encryptconn.New(conn)
			go r.handle(clearConn)
		}
		r.log.Info().Msg("stopped " + r.addr)
	}()
	<-started
	r.log.Info().Msg("started " + r.addr)
	return nil
}

func (r *rx) handle(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	data := make([]byte, 4096)
	i, err := conn.Read(data)
	if err != nil {
		r.log.Debug().Err(err).Msgf("read err: %v", data)
		return
	}
	msg := &hbtype.Msg{}
	if err := json.Unmarshal(data[:i], msg); err != nil {
		r.log.Error().Err(err).Msgf("Unmarshal data: %s", data[:i])
		return
	}
	r.cmdC <- hbctrl.CmdSetPeerSuccess{
		Nodename: msg.Nodename,
		HbId:     r.id,
		Success:  true,
	}
	r.msgC <- msg
}

func newRx(ctx context.Context, name string, nodes []string, addr, port, intf string, timeout time.Duration) *rx {
	id := name + ".rx"
	log := daemonctx.Logger(ctx).With().Str("id", id).Logger()
	return &rx{
		ctx:     ctx,
		id:      id,
		nodes:   nodes,
		addr:    addr,
		port:    port,
		intf:    intf,
		timeout: timeout,
		log:     log,
	}
}