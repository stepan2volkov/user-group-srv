PHONY: .generate
.generate:
	 go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	 cp ${GOPATH}/bin/oapi-codegen ./bin/oapi-codegen

PHONY: generate
generate: .generate
	./bin/oapi-codegen -config openapi-generate.yaml ./internal/api/openapi/openapi.yaml
	