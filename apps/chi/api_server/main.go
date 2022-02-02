package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/jirokun/go_frameworks/chi/handler"
)

func main() {
	swagger, err := handler.GetSwagger()
	if err != nil {
		fmt.Printf("failed to get swagger spec: %v\n", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	openapi3.DefineIPv4Format()
	openapi3.DefineIPv6Format()

	r := chi.NewRouter()
	si := InitializeServerInterface()
	handler.HandlerFromMux(si, r)
	http.ListenAndServe(":3000", r)
}
