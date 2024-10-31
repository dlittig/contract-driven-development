// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all pets
	// (GET /pets)
	ListPets(c *gin.Context, params ListPetsParams)
	// Create a pet
	// (POST /pets)
	CreatePets(c *gin.Context)
	// Info for a specific pet
	// (GET /pets/{petId})
	ShowPetById(c *gin.Context, petId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ListPets operation middleware
func (siw *ServerInterfaceWrapper) ListPets(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPetsParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListPets(c, params)
}

// CreatePets operation middleware
func (siw *ServerInterfaceWrapper) CreatePets(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreatePets(c)
}

// ShowPetById operation middleware
func (siw *ServerInterfaceWrapper) ShowPetById(c *gin.Context) {

	var err error

	// ------------- Path parameter "petId" -------------
	var petId string

	err = runtime.BindStyledParameterWithOptions("simple", "petId", c.Param("petId"), &petId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter petId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ShowPetById(c, petId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/pets", wrapper.ListPets)
	router.POST(options.BaseURL+"/pets", wrapper.CreatePets)
	router.GET(options.BaseURL+"/pets/:petId", wrapper.ShowPetById)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xVTW/bMAz9KwK3wwZ4sdPefFuHAguwDQXaW5GDZtOOOuujEp3GCPzfB0pO0jbpeumA",
	"niLJFPne42O0hcpqZw0aClBuIVQr1DIuL723nhfOW4eeFMbjytbIv431WhKUoAydn0EGNDhMW2zRw5iB",
	"xhBkG6Onj4G8Mi2MYwYe73vlsYbyNuU8xC/3yezvO6yIc10hHWNR9YncGRip8eQHku3rYFQNU4oXcKTS",
	"hDouPnpsoIQP+UHIfFIxZ9Csg9wsUvi8KPY5pfdygJHLK9NYztWpCk2I2BMH+Lm4icAVdby9fpBti14w",
	"CrKeRVujD8oaKGE+K2YFR1uHRjoFJZzHowycpFVEm7sJf5v0ZDUlKWsWNZTwQwWKBPmGlxoJfYDydgs1",
	"hsorR6nSd/sgtDSDiCoIssIj9d4IScIaFKQ0ik9absS8KD4DE4QS7nv0w07bEjqlFUE2We6kpbTcKN3r",
	"p7rtDTYuuXHBWROSHc6KIjnUEJpITzrXqSoSzO8CY98+qvdK60JqzlPqX4WTLdYidk/YRrik1wplHcXa",
	"wuaLwU0s//xqp8wfVotWKDgm5uIkBxqP9Xju04SmkX1Hb8YzTfkJor3BjcOKsBY4xWQQeq2lHyarCNl1",
	"O/4kW3YKxO1yzMDZcMJh3zxKwsljPHUY6MLWw1v2LbE5jDT5Hscjq8yPG/Sr77p9K+AdiZ1UE5LFPtZ6",
	"zNJY51uHtKjHF8f7emUfrpAuhkX92oTfrFComq3JXnVI05B7hWvcDTT/qxzmORaH58L/y8//eX5PCX25",
	"k3lXmHlJsZad4rPox/fU+YVprGisF1IEh5VqVPWCCfga+vWunb3voIQVkSvz6A5+LmYhvR8zZfP1HMbl",
	"+DcAAP//rLWv9AAIAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
