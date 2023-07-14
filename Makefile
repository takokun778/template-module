include .env
export

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: aqua
aqua: ## Put the path in your environment variables. ex) export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
	@go run github.com/aquaproj/aqua-installer@latest --aqua-version v2.0.0

.PHONY: tool
tool: ## install tool
	@aqua install
	@(cd api && npm install)

.PHONY: goinit
goinit: ## module init
	@rm -rf go.mod go.sum
	@rm -rf vendor
	@go mod init ${REPOSITORY}
	@go mod tidy
	@go mod vendor

.PHONY: gocompile
gocompile: ## go compile
	@go build -v ./... && go clean

.PHONY: gofmt
gofmt: ## go format
	@go fmt ./...

.PHONY: golint
golint: ## go lint
	@golangci-lint run --fix

.PHONY: gomod
gomod: ## go mod tidy & go mod vendor
	@go mod tidy
	@go mod vendor

.PHONY: gorenovate
gorenovate: ## go modules update
	@go get -u -t ./...
	@go mod tidy
	@go mod vendor

.PHONY: gogen
gogen: ## Generate code.
	@go generate ./...
	@oapi-codegen -generate types -package openapi ./api/openapi.yaml > ./pkg/openapi/types.gen.go
	@oapi-codegen -generate chi-server -package openapi ./api/openapi.yaml > ./pkg/openapi/server.gen.go
	@oapi-codegen -generate client -package openapi ./api/openapi.yaml > ./pkg/openapi/client.gen.go
	@(cd proto && buf generate --template buf.gen.yaml)
	@go mod tidy
	@go mod vendor

.PHONY: gotest
gotest: ## unit test
	@$(call _test,${c})

define _test
if [ -z "$1" ]; then \
	go test ./internal/... ; \
else \
	go test ./internal/... -count=1 ; \
fi
endef

.PHONY: e2e
e2e: ## e2e test
	@$(call _e2e,${c})

define _e2e
if [ -z "$1" ]; then \
	go test ./e2e/... ; \
else \
	go test ./e2e/... -count=1 ; \
fi
endef

.PHONY: up
up: ## docker compose up with air hot reload
	@docker compose --project-name ${APP_NAME} --file ./.docker/compose.yaml up -d

.PHONY: down
down: ## docker compose down
	@docker compose --project-name ${APP_NAME} down --volumes

.PHONY: balus
balus: ## Destroy everything about docker. (containers, images, volumes, networks.)
	@docker compose --project-name ${APP_NAME} down --rmi all --volumes

.PHONY: ymlfmt
ymlfmt: ## Format yaml file
	@yamlfmt

.PHONY: ymlint
ymlint: ## Lint yaml file
	@yamlfmt -lint

.PHONY: apilint
apilint: ## Lint api file
	@(cd api && npx spectral lint openapi.yaml)
