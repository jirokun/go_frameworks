# note: call scripts from /scripts
.PHONY: chi-gen-openapi
chi-oapi-codegen-gen-openapi:
	oapi-codegen -package handler -generate chi-server,spec api/petstore-v3.yaml > apps/chi/oapi-codegen/handler/handler.gen.go
	oapi-codegen -package usecase -generate types api/petstore-v3.yaml > apps/chi/oapi-codegen/usecase/types.gen.go

.PHONY: chi-wire
chi-oapi-codegen-wire:
	cd apps/chi/oapi-codegen && $$GO_PATH/bin/wire

.PHONY: chi-run 
chi-oapi-codegen-run:
	cd apps/chi/oapi-codegen && go run .



.PHONY: chi-quicktype-gen-quicktype
chi-quicktype-gen-quicktype:
	quicktype -o apps/chi/quicktype/usecase/input/new_pet.gen.go --src-lang typescript --lang go --package input api/quicktype/input/NewPet.ts
	quicktype -o apps/chi/quicktype/usecase/output/pet.gen.go --src-lang typescript --lang go --package output api/quicktype/output/Pet.ts

.PHONY: chi-quicktype-wire
chi-quicktype-wire:
	cd apps/chi/quicktype && $$GO_PATH/bin/wire

.PHONY: chi-quicktype-run 
chi-quicktype-run:
	cd apps/chi/quicktype && go run .
