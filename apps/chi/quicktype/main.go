package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	petHandler := InitializePetHandler()
	r.Mount("/", petHandler)
	http.ListenAndServe(":3000", r)
}
