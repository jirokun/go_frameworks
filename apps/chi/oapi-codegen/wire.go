// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jirokun/go_frameworks/apps/chi/oapi-codegen/handler"
	"github.com/jirokun/go_frameworks/apps/chi/oapi-codegen/infra"
	"github.com/jirokun/go_frameworks/apps/chi/oapi-codegen/usecase"
)

func InitializeServerInterface() *handler.ServerInterfaceImpl {
	wire.Build(
		handler.NewServerInterfaceImpl,
		usecase.NewPetUsecase,
		infra.NewPetRepositoryImpl,
	)
	return &handler.ServerInterfaceImpl{}
}
