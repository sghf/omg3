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

	// (GET /openapi)
	GetSwagger(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

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
		r.Get(options.BaseURL+"/openapi", wrapper.GetSwagger)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"nx0eDv8OAAD//x5k0Zy9KQAA",
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