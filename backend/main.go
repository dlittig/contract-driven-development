package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	api "cdd/backend/api/petstore"
	gen "cdd/backend/api/petstore/gen"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	middleware "github.com/oapi-codegen/gin-middleware"
)

func NewGinPetServer(petStore *api.PetStore, port string) *http.Server {
	swagger, err := gen.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic gin router
	r := gin.Default()

	// Add CORS middleware
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	gen.RegisterHandlers(r, petStore)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}
	return s
}

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface
	petStore := api.NewPetStore()
	s := NewGinPetServer(petStore, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
