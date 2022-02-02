# note: call scripts from /scripts
.PHONY: chi-gen-openapi
chi-gen-openapi:
	oapi-codegen -package handler -generate chi-server,spec api/petstore-expanded.yaml > apps/chi/api_server/handler/handler.gen.go
	oapi-codegen -package usecase -generate types api/petstore-expanded.yaml > apps/chi/api_server/usecase/types.gen.go

.PHONY: chi-wire
chi-wire:
	cd apps/chi/api_server && $$GO_PATH/bin/wire

.PHONY: chi-run 
chi-run:
	cd apps/chi/api_server && go run .