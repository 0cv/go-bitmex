DEFAULT_GOAL := default

default: swagger

swagger:
	@scripts/swagger.sh

install-swagger:
	@curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger

.PHONY: swagger install-swagger