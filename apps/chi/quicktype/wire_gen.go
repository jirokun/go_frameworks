// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/handler"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/infra"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase"
)

// Injectors from wire.go:

func InitializePetHandler() *handler.PetHandler {
	petRepository := infra.NewPetRepositoryImpl()
	petUsecase := usecase.NewPetUsecase(petRepository)
	petHandler := handler.NewPetHandler(petUsecase)
	return petHandler
}
