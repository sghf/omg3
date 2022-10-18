// Package daemonapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package daemonapi

import (
	"bytes"
	"compress/gzip"
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

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectStatus(w, r)
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

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xaX4/buBH/KoRaoHeAYu/mT4H6qdfDtQjQNkG3b9lFMBbHEu8oUiEp77oHf/diSFGS",
	"JWptt9midy+JV+T8Hw7nN9LPWaHrRitUzmabn7MGDNTo0Pi/vrRoDh+2P2LhPoKr6BFHWxjROKFVtsm0",
	"X2MNLeaZoEeeJsszBTVmm6xbMvilFQZ5tnGmxTyzRYU1EEN3aGifdUaoMjse87HUO5RYOG0WJdu4IS19",
	"tHy5Bse46H0ATZPYlGdoTNCrMbpB4wT67YXmONf2B9rM/Fqe7bSpwWWbTCj35nWWR95COSzREPMarYVy",
	"kVFczhPeG+z8lHUC4/aHI3nJOlAF3jlwrZ3r35n7W4O7bJP9Zj2kx7pzypq2HPMM9iDkub02iDnmWVEJ",
	"yQ2qcxSUMP9ACd5gotPKOgOiS9DO4K3WEiFssG2djBA3TZoC1T5JsJP49LmGp9HiKCZhVahnVh2YEt3C",
	"BqP/Fazv48/B4Ssn6kQg8+wnofg5X/k9xzzTPj3gimhoU1RIfnV4jmi8lSj3aEBeIaoBE6vLNXFvJBRY",
	"o3JnCfuNRGXQotkjD0dnB6102WYH0mI+OUpxKxOWUUlgYsdcJSwLqrMKLFPasS2iYm1DweKMt8icZsDu",
	"VYVg3BbBMa4fFYWRFeQc5Gx7YMBqyllUdNhYg0ZovrpXjxUq5ipMrDJU3OZ+sdPAVrqVnG2RtaqoQJXI",
	"c3avQHHWK/8opKQdFh0p5i1d3asho0Z53xihjXCHsx6N+zyN3gsrtEJ+nmzY6guR1a0pQlkRDuuzKRAp",
	"fnhqtEV+16dQZwoYA14p0ypFx2TMeHaApkS2AInps2kl7PHqDA1R+lwa3aavCNtuLYbMP0293jVsVHzz",
	"wZbTkkx3spQoUyk9D7IRPH2tji8G2pQPjB96RuFaTbnP6UZLXZ5Nnn7fMc+6U3Np0ZsoGS6YvnJ2JXGo",
	"QKfJOUhLWROr6SxGSnN8r3Z67nYJW5SJ4IXnvmpUyKSwjukdIz4sLK3GoXzOVUTzVyJJ+ZsWQxczVSCu",
	"RBX8b70Lv0mNxwoNBu2Crr5igKssA0NVqhaqZDuj61Xq5vE752IDg5TZTjPrtIESmVefWVBB3sWusKB8",
	"hzlzxCQneq/kMT5R31TUBwfPorvg2pFbPX/v3KSX9iDbBAf/+JSFf7Q6m+6dWYHvkjU25urFCeYJEvkV",
	"+H6v1U6Uc/dwcOPeeFCi9kf38gM9YxB+/VmEWpyW2vPeHlyyObpWi7Gf626XF/awqGEEPDPZeoZMLorF",
	"iGsqGqf92LTMo6Ie91NWQebvHeOy3C89LJzf/p4aq7eI3Zg2DBSL2CA8+4b+/SOl0LepEEwtOOnXev0z",
	"pZV3dlJ0JGGNlqIg3BYN7ch8yjNtOBKEkxo4g31JPqjEztH/jUGg4m8LbcbHZuQObV1Ak99ttXGp+zUZ",
	"6Una+F2pdBn4fy8RzAvy/5tWwqVAZyn1FuRnfGqIJpW0X02HJdzYLB0X2xM8dzomsDSpWM8rqeGot+2z",
	"793NNPNIHm8l3X+Rgq4xUKzLyf4MaMWACWf9FZcE6ZPWeCIIzV4UmI8YGhb7PjYi9c0+DqnvpxPdsa/F",
	"k+9s1BqSqZ1umOdlVVjYysRlVRGq9jdVr5kolTZoGUgZNGPOgLKCKBgU9J9NggtUBTRzEUJxUYBDEgNu",
	"IotAluIyICZa8kxsKz3WgpJSJCKeoBhnHZPq0JCHrTZM4h7lAuQR3XV5qtRPeHgVLuoGhLEhHJySgsJr",
	"0LrwO9wuZLnTrNCSSj67J2/gq0fBkcFWty6gxmjVWJEhPWXsQmYxlPoUx0yaULItSTcaET1/qAOL0RTo",
	"gma/HirNqeccShkyhgPWWhFoFi4iVWdEWaIh8BsYdBnDeth7r8bRJ3jdNguhGw81Jjk1eDtCZShLg6VP",
	"G6GcZh8CRvCnD4HTGf+O4MRwHAPh6l75kZplQrEoceDOtfqdo+a2YbB0HBbB9sXAOYr7GEkG5Eu5CGZh",
	"pNRBvUtYv+ddLVZ8e1gGpDGQIB/hYP3kockZ7lEx2DkfWe+M61xx2RUwTIwCbk4k3wTshH2nx4/SCqwV",
	"JZVWp5OdC5T2utFB+PsyUB1Oem90R5w6dKPYpJRYyopZjbiqGQ6K4VWznImdgUHKoojiZioKJZyArqJc",
	"gAPf9/t9vOJU9QLKf4bNU50HBXp+zxjwfqzu9D7rliLGq7R1zNJlEFEvQ8UbLZRbhRHFxagT2KM2kvub",
	"pVXiS4un/JjgqJzYCTTEGp+gbqSvCV/U6vXNzdtXtzerQterdtsq125ubjf4+y1/C2+27969XW7iZ+fs",
	"0PQQtpdNDydSbWHFZaDvNDhzgf55FDmZJfxfuPYPr25vvWt1g8rui5U1+w3H/Wt1u+r0XQUrVrfXOxq+",
	"pqv7KjFXYjTCGxDaDoTUew+wkiitpxrQ2YhkJ/Ep0ZseR72XE84b0bmOQUOm7NHYIOfN6s3qNtz4qGiR",
	"Ht2sbrLRKMqPMggmBJZdFlH8fX9GJTT7C7q/9yMSXz8brWzIjtc3N+HtnHIdRIWmkdSZCq3WP9qAlId3",
	"guemKkGIN3PS97dFgZYQDK11qH8NPfLUNqH5xwlEDbULrfuT5oevpvcUCCe09wtsGEbQ1V60xqBy8sC6",
	"sXuARuHVKm3Quw48+XHf6XvWYzoOCz7Lh7T8SjaHt7QJS1sV0DJy1u8ZBazoofyZgAXM/9IBC1ISZvgF",
	"G0yw7JudNqyrGjnznSOjs4r8W+pv/XunrjH3l7j12PfXHsp+xLlUNT6MR6H5yTcQn9I6DVvW028kjvnU",
	"J0hQM2D81DcKo+XZNwl9Ez3nKuoGjdXKN+J0ewQ2vhnvhxYpeSPC7LnPIB5esISeDJ8vrqK7bmj8fBz9",
	"aPm/juLLW+/1vNj2ESI/V5LimPCli1KUk7ChPlHhl19D7Gja/3z23Q1fHP3HGdjz+B9k4SDr0kwcGsxz",
	"iXgXIfDL5mE/M54ZQHviYNeOlfllZmQcXXzyQyCjQGYP4XM1NPuYZq2R2SarnGs267XUBUhCqJvb17fv",
	"1tRcHx+O/w4AAP//ywGqDPUnAAA=",
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
