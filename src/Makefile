GOROOT ?= /usr/local/go
GOPATH := $(shell go env GOPATH)

# need v1.54.2 of golangci-lint
lint:
	golangci-lint run -v -c golangci.yml ./...

migration-create-sql:
	goose -dir=./migrations create $(NAME) sql