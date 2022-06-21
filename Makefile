BUILD_COMMIT := $(shell git log --format="%H" -n 1)
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%S')

APP_NAME = user-group-srv
PROJECT = github.com/stepan2volkov/$(APP_NAME)
CMD = $(PROJECT)/cmd/$(APP_NAME)

PHONY: .generate
.generate:
	 go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	 go install github.com/pressly/goose/v3/cmd/goose@latest
	 cp ${GOPATH}/bin/oapi-codegen ./bin/oapi-codegen
	 cp ${GOPATH}/bin/goose ./bin/goose

PHONY: generate
generate: .generate
	./bin/oapi-codegen -config openapi-generate.yaml ./internal/api/openapi/openapi.yaml

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="\
		-X '$(PROJECT)/internal/config.BuildCommit=$(BUILD_COMMIT)'\
		-X '${PROJECT}/internal/config.BuildTime=${BUILD_TIME}'"\
		-o build $(CMD)

DSN := "user=test port=5432 password=1234 dbname=usergroup sslmode=disable"
migrate:
	goose -dir=./migrations postgres $(DSN) up

build_container:
	docker build -t docker.io/stepan2volkov/$(APP_NAME) -f ./build/Dockerfile .

push_container:
	docker push docker.io/stepan2volkov/${APP_NAME}