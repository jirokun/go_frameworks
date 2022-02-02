# note: call scripts from /scripts
.PHONY: gen-openapi
gen-openapi:
	oapi-codegen -package handler -generate chi-server api/petstore-expanded.yaml > apps/chi/api_server/handler/handler.gen.go

.PHONY: wire
wire:
	cd apps/chi/api_server && $$GO_PATH/bin/wire
