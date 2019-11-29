DEFAULT_GOAL := default

default: lint test

test: lint
	@go test -v

swagger:
	@scripts/swagger.sh

lint: install-golangci-lint
	@golangci-lint run --timeout 3m0s

install-swagger:
	@curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger

install-golangci-lint:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: swagger install-swagger test lint install-golangci-lint