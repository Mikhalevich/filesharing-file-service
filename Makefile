all: build

.PHONY: build
build:
	go build -mod=vendor -o ./bin/file cmd/file/main.go

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
