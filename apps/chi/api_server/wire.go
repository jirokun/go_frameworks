// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jirokun/go_frameworks/chi/handler"
	"github.com/jirokun/go_frameworks/chi/infra"
	"github.com/jirokun/go_frameworks/chi/usecase"
)

func InitializeServerInterface() *handler.ServerInterfaceImpl {
	wire.Build(
		handler.NewServerInterfaceImpl,
		usecase.NewPetUsecase,
		infra.NewPetRepositoryImpl,
	)
	return &handler.ServerInterfaceImpl{}
}
