.PHONY: generate
generate:
	@echo "Generating..."
	@oapi-codegen -generate types,client -package api schema.yaml > api/api.gen.go

.PHONY: build
build: generate
	@echo "Building..."
	@goreleaser build --snapshot
