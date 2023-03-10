/*
Package rawmux provides raw multiplexer from httpmux

It can be used by raw listeners to Serve accepted connexions
*/
package routeraw

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	clientrequest "github.com/opensvc/om3/core/client/request"
	"github.com/opensvc/om3/daemon/ccfg"
	"github.com/opensvc/om3/daemon/daemonenv"
	"github.com/opensvc/om3/daemon/listener/routeresponse"
)

type (
	T struct {
		httpMux http.Handler
		log     zerolog.Logger
		timeOut time.Duration
	}

	ReadWriteCloseSetDeadliner interface {
		io.ReadWriteCloser
		SetDeadline(time.Time) error
		SetWriteDeadline(time.Time) error
	}

	srcNoder interface {
		SrcNode() string
	}

	// request struct holds the translated raw request for http mux
	request struct {
		method  string
		path    string
		handler http.HandlerFunc
		body    io.Reader
		header  http.Header
	}
)

// New function returns an initialised *T that will use mux as http mux
func New(mux http.Handler, log zerolog.Logger, timeout time.Duration) *T {
	return &T{
		httpMux: mux,
		log:     log,
		timeOut: timeout,
	}
}

// Serve function is an adapter to serve raw call from http mux
//
// # Serve can be used on raw listeners accepted connexions
//
// 1- raw request will be decoded to create to http request
// 2- http request will be served from http mux ServeHTTP
// 3- Response is sent to w
func (t *T) Serve(rw ReadWriteCloseSetDeadliner) {
	defer func() {
		err := rw.Close()
		if err != nil {
			t.log.Debug().Err(err).Msg("rawunix.Serve close failure")
			return
		}
	}()
	// TODO some handlers needs no deadline
	//if err := rw.SetWriteDeadline(time.Now().Add(t.timeOut)); err != nil {
	//	t.log.Error().Err(err).Msg("rawunix.Serve can't set SetDeadline")
	//}
	req, err := t.newRequestFrom(rw)
	if err != nil {
		// Don't warn on empty early closed rw, it may be arbitrator tcp dial check
		if !errors.Is(err, io.EOF) {
			t.log.Warn().Err(err).Msg("rawunix.Serve can't analyse request")
		} else {
			t.log.Debug().Err(err).Msg("rawunix.Serve can't analyse request")
		}
		return
	}
	resp := routeresponse.NewResponse(rw)
	if err := req.do(resp); err != nil {
		t.log.Error().Err(err).Msgf("rawunix.Serve request.do error for %s %s",
			req.method, req.path)
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.log.Error().Msgf("rawunix.Serve unexpected status code %d for %s %s",
			resp.StatusCode, req.method, req.path)
		return
	}
	t.log.Info().Msgf("status code is %d", resp.StatusCode)
}

// newRequestFrom functions returns *request from w
func (t *T) newRequestFrom(rw io.ReadWriteCloser) (*request, error) {
	var b = make([]byte, 4096)
	_, err := rw.Read(b)
	if err != nil {
		t.log.Debug().Err(err).Msg("newRequestFrom read failure")
		return nil, err
	}
	srcRequest := clientrequest.T{}
	b = bytes.TrimRight(b, "\x00")
	if err := json.Unmarshal(b, &srcRequest); err != nil {
		t.log.Warn().Err(err).Msgf("newRequestFrom invalid message: %s", string(b))
		return nil, err
	}
	t.log.Debug().Msgf("newRequestFrom: %s, options: %s", srcRequest, srcRequest.Options)
	matched, ok := actionToPath[srcRequest.Action]
	if !ok {
		matched.method = srcRequest.Method
		matched.path = srcRequest.Action
	}
	reqUrl := url.URL{
		Path:     matched.path,
		RawQuery: srcRequest.Values.Encode(),
	}
	httpHeader := http.Header{}
	if srcRequest.Node != "" {
		httpHeader.Set(daemonenv.HeaderNode, srcRequest.Node)
	} else if noder, ok := rw.(srcNoder); ok {
		httpHeader.Set(daemonenv.HeaderNode, noder.SrcNode())
	}
	return &request{
		method:  matched.method,
		path:    reqUrl.RequestURI(),
		handler: t.httpMux.ServeHTTP,
		body:    bytes.NewReader(b),
		header:  httpHeader,
	}, nil
}

// do function execute http mux handler on translated request and returns error
func (r *request) do(resp *routeresponse.Response) error {
	body := r.body
	request, err := http.NewRequest(r.method, r.path, body)
	request.Header = r.header
	request.SetBasicAuth(r.header.Get(daemonenv.HeaderNode), ccfg.Get().Secret())
	if err != nil {
		return err
	}
	r.handler(resp, request)
	return nil
}
