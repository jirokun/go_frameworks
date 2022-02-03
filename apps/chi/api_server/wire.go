// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jirokun/go_frameworks/apps/chi/api_server/handler"
	"github.com/jirokun/go_frameworks/apps/chi/api_server/infra"
	"github.com/jirokun/go_frameworks/apps/chi/api_server/usecase"
)

func InitializeServerInterface() *handler.ServerInterfaceImpl {
	wire.Build(
		handler.NewServerInterfaceImpl,
		usecase.NewPetUsecase,
		infra.NewPetRepositoryImpl,
	)
	return &handler.ServerInterfaceImpl{}
}
