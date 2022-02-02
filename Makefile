# note: call scripts from /scripts
.PHONY: gen-openapi
gen-openapi:
	oapi-codegen -package handler -generate chi-server,spec api/petstore-expanded.yaml > apps/chi/api_server/handler/handler.gen.go
	oapi-codegen -package usecase -generate types api/petstore-expanded.yaml > apps/chi/api_server/usecase/types.gen.go

.PHONY: wire
wire:
	cd apps/chi/api_server && $$GO_PATH/bin/wire

.PHONY: run 
run:
	cd apps/chi/api_server && go run .