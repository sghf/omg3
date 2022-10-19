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

	// (POST /daemon/stop)
	PostDaemonStop(w http.ResponseWriter, r *http.Request)

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

	// (GET /object/selector)
	GetObjectSelector(w http.ResponseWriter, r *http.Request, params GetObjectSelectorParams)

	// (POST /object/status)
	PostObjectStatus(w http.ResponseWriter, r *http.Request)

	// (GET /public/openapi)
	GetSwagger(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// PostDaemonStop operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonStop(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonStop(w, r)
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

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectMonitor(w, r)
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

// PostObjectStatus operation middleware
func (siw *ServerInterfaceWrapper) PostObjectStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectStatus(w, r)
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
		r.Post(options.BaseURL+"/daemon/stop", wrapper.PostDaemonStop)
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
		r.Get(options.BaseURL+"/object/selector", wrapper.GetObjectSelector)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/status", wrapper.PostObjectStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/public/openapi", wrapper.GetSwagger)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"IdMqAAA=",
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
