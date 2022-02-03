// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/handler"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/infra"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase"
)

func InitializePetHandler() *handler.PetHandler {
	wire.Build(
		handler.NewPetHandler,
		usecase.NewPetUsecase,
		infra.NewPetRepositoryImpl,
	)
	return &handler.PetHandler{}
}
