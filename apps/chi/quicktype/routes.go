package main

import (
	"github.com/go-chi/chi/v5"
)

func routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Route("/pets/", func(r chi.Router) {
			petHandler := InitializePetHandler()
			r.Get("/{id}", petHandler.FindPetById)
			r.Post("/", petHandler.AddPet)
		})
	})

	return r
}
