// Package daemonapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package daemonapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/token)
	PostAuthToken(w http.ResponseWriter, r *http.Request, params PostAuthTokenParams)

	// (GET /daemon/dns/dump)
	GetDaemonDNSDump(w http.ResponseWriter, r *http.Request)

	// (GET /daemon/events)
	GetDaemonEvents(w http.ResponseWriter, r *http.Request, params GetDaemonEventsParams)

	// (POST /daemon/join)
	PostDaemonJoin(w http.ResponseWriter, r *http.Request, params PostDaemonJoinParams)

	// (POST /daemon/leave)
	PostDaemonLeave(w http.ResponseWriter, r *http.Request, params PostDaemonLeaveParams)

	// (POST /daemon/logs/control)
	PostDaemonLogsControl(w http.ResponseWriter, r *http.Request)

	// (GET /daemon/running)
	GetDaemonRunning(w http.ResponseWriter, r *http.Request)

	// (GET /daemon/status)
	GetDaemonStatus(w http.ResponseWriter, r *http.Request, params GetDaemonStatusParams)

	// (POST /daemon/stop)
	PostDaemonStop(w http.ResponseWriter, r *http.Request)

	// (POST /daemon/sub/action)
	PostDaemonSubAction(w http.ResponseWriter, r *http.Request)

	// (POST /instance/status)
	PostInstanceStatus(w http.ResponseWriter, r *http.Request)

	// (POST /node/clear)
	PostNodeClear(w http.ResponseWriter, r *http.Request)

	// (GET /node/file)
	GetNodeFile(w http.ResponseWriter, r *http.Request, params GetNodeFileParams)

	// (POST /node/monitor)
	PostNodeMonitor(w http.ResponseWriter, r *http.Request)

	// (GET /nodes/info)
	GetNodesInfo(w http.ResponseWriter, r *http.Request)

	// (POST /object/abort)
	PostObjectAbort(w http.ResponseWriter, r *http.Request)

	// (POST /object/clear)
	PostObjectClear(w http.ResponseWriter, r *http.Request)

	// (GET /object/config)
	GetObjectConfig(w http.ResponseWriter, r *http.Request, params GetObjectConfigParams)

	// (GET /object/file)
	GetObjectFile(w http.ResponseWriter, r *http.Request, params GetObjectFileParams)

	// (POST /object/monitor)
	PostObjectMonitor(w http.ResponseWriter, r *http.Request)

	// (POST /object/progress)
	PostObjectProgress(w http.ResponseWriter, r *http.Request)

	// (GET /object/selector)
	GetObjectSelector(w http.ResponseWriter, r *http.Request, params GetObjectSelectorParams)

	// (POST /object/switchTo)
	PostObjectSwitchTo(w http.ResponseWriter, r *http.Request)

	// (GET /public/openapi)
	GetSwagger(w http.ResponseWriter, r *http.Request)

	// (GET /relay/message)
	GetRelayMessage(w http.ResponseWriter, r *http.Request, params GetRelayMessageParams)

	// (POST /relay/message)
	PostRelayMessage(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// PostAuthToken operation middleware
func (siw *ServerInterfaceWrapper) PostAuthToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostAuthTokenParams

	// ------------- Optional query parameter "role" -------------
	if paramValue := r.URL.Query().Get("role"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "role", r.URL.Query(), &params.Role)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "role", Err: err})
		return
	}

	// ------------- Optional query parameter "duration" -------------
	if paramValue := r.URL.Query().Get("duration"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "duration", r.URL.Query(), &params.Duration)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "duration", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthToken(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDaemonDNSDump operation middleware
func (siw *ServerInterfaceWrapper) GetDaemonDNSDump(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDaemonDNSDump(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDaemonEvents operation middleware
func (siw *ServerInterfaceWrapper) GetDaemonEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonEventsParams

	// ------------- Optional query parameter "duration" -------------
	if paramValue := r.URL.Query().Get("duration"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "duration", r.URL.Query(), &params.Duration)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "duration", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "filter" -------------
	if paramValue := r.URL.Query().Get("filter"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "filter", r.URL.Query(), &params.Filter)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filter", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDaemonEvents(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostDaemonJoin operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonJoin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostDaemonJoinParams

	// ------------- Required query parameter "node" -------------
	if paramValue := r.URL.Query().Get("node"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "node"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "node", r.URL.Query(), &params.Node)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "node", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonJoin(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostDaemonLeave operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonLeave(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostDaemonLeaveParams

	// ------------- Required query parameter "node" -------------
	if paramValue := r.URL.Query().Get("node"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "node"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "node", r.URL.Query(), &params.Node)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "node", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonLeave(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostDaemonLogsControl operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonLogsControl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonLogsControl(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDaemonRunning operation middleware
func (siw *ServerInterfaceWrapper) GetDaemonRunning(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDaemonRunning(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDaemonStatus operation middleware
func (siw *ServerInterfaceWrapper) GetDaemonStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonStatusParams

	// ------------- Optional query parameter "namespace" -------------
	if paramValue := r.URL.Query().Get("namespace"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	// ------------- Optional query parameter "relatives" -------------
	if paramValue := r.URL.Query().Get("relatives"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "relatives", r.URL.Query(), &params.Relatives)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "relatives", Err: err})
		return
	}

	// ------------- Optional query parameter "selector" -------------
	if paramValue := r.URL.Query().Get("selector"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "selector", r.URL.Query(), &params.Selector)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "selector", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDaemonStatus(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostDaemonStop operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonStop(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonStop(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostDaemonSubAction operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonSubAction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonSubAction(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostInstanceStatus operation middleware
func (siw *ServerInterfaceWrapper) PostInstanceStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostInstanceStatus(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostNodeClear operation middleware
func (siw *ServerInterfaceWrapper) PostNodeClear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNodeClear(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodeFile operation middleware
func (siw *ServerInterfaceWrapper) GetNodeFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodeFileParams

	// ------------- Required query parameter "kind" -------------
	if paramValue := r.URL.Query().Get("kind"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "kind"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "kind", r.URL.Query(), &params.Kind)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "kind", Err: err})
		return
	}

	// ------------- Required query parameter "name" -------------
	if paramValue := r.URL.Query().Get("name"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "name"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "name", r.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodeFile(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostNodeMonitor operation middleware
func (siw *ServerInterfaceWrapper) PostNodeMonitor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNodeMonitor(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodesInfo operation middleware
func (siw *ServerInterfaceWrapper) GetNodesInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodesInfo(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectAbort operation middleware
func (siw *ServerInterfaceWrapper) PostObjectAbort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectAbort(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectClear operation middleware
func (siw *ServerInterfaceWrapper) PostObjectClear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectClear(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetObjectConfig operation middleware
func (siw *ServerInterfaceWrapper) GetObjectConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectConfigParams

	// ------------- Required query parameter "path" -------------
	if paramValue := r.URL.Query().Get("path"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "path"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "path", r.URL.Query(), &params.Path)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	// ------------- Optional query parameter "evaluate" -------------
	if paramValue := r.URL.Query().Get("evaluate"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "evaluate", r.URL.Query(), &params.Evaluate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "evaluate", Err: err})
		return
	}

	// ------------- Optional query parameter "impersonate" -------------
	if paramValue := r.URL.Query().Get("impersonate"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "impersonate", r.URL.Query(), &params.Impersonate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "impersonate", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetObjectConfig(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetObjectFile operation middleware
func (siw *ServerInterfaceWrapper) GetObjectFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectFileParams

	// ------------- Required query parameter "path" -------------
	if paramValue := r.URL.Query().Get("path"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "path"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "path", r.URL.Query(), &params.Path)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetObjectFile(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectMonitor operation middleware
func (siw *ServerInterfaceWrapper) PostObjectMonitor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectMonitor(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectProgress operation middleware
func (siw *ServerInterfaceWrapper) PostObjectProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectProgress(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetObjectSelector operation middleware
func (siw *ServerInterfaceWrapper) GetObjectSelector(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectSelectorParams

	// ------------- Required query parameter "selector" -------------
	if paramValue := r.URL.Query().Get("selector"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "selector"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "selector", r.URL.Query(), &params.Selector)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "selector", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetObjectSelector(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectSwitchTo operation middleware
func (siw *ServerInterfaceWrapper) PostObjectSwitchTo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectSwitchTo(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetSwagger operation middleware
func (siw *ServerInterfaceWrapper) GetSwagger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSwagger(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetRelayMessage operation middleware
func (siw *ServerInterfaceWrapper) GetRelayMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRelayMessageParams

	// ------------- Optional query parameter "nodename" -------------
	if paramValue := r.URL.Query().Get("nodename"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "nodename", r.URL.Query(), &params.Nodename)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "nodename", Err: err})
		return
	}

	// ------------- Optional query parameter "cluster_id" -------------
	if paramValue := r.URL.Query().Get("cluster_id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "cluster_id", r.URL.Query(), &params.ClusterId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cluster_id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRelayMessage(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostRelayMessage operation middleware
func (siw *ServerInterfaceWrapper) PostRelayMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRelayMessage(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/token", wrapper.PostAuthToken)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/daemon/dns/dump", wrapper.GetDaemonDNSDump)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/daemon/events", wrapper.GetDaemonEvents)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/daemon/join", wrapper.PostDaemonJoin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/daemon/leave", wrapper.PostDaemonLeave)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/daemon/logs/control", wrapper.PostDaemonLogsControl)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/daemon/running", wrapper.GetDaemonRunning)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/daemon/status", wrapper.GetDaemonStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/daemon/stop", wrapper.PostDaemonStop)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/daemon/sub/action", wrapper.PostDaemonSubAction)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/instance/status", wrapper.PostInstanceStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/node/clear", wrapper.PostNodeClear)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/node/file", wrapper.GetNodeFile)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/node/monitor", wrapper.PostNodeMonitor)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/nodes/info", wrapper.GetNodesInfo)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/abort", wrapper.PostObjectAbort)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/clear", wrapper.PostObjectClear)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/object/config", wrapper.GetObjectConfig)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/object/file", wrapper.GetObjectFile)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/monitor", wrapper.PostObjectMonitor)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/progress", wrapper.PostObjectProgress)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/object/selector", wrapper.GetObjectSelector)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/switchTo", wrapper.PostObjectSwitchTo)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/public/openapi", wrapper.GetSwagger)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/relay/message", wrapper.GetRelayMessage)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/relay/message", wrapper.PostRelayMessage)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
