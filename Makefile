.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt # ignores any directory named fmt

lint: fmt
	golangci-lint run
.PHONY:lint # ignores any directory named lint

vet: fmt
	go vet ./...
.PHONY:vet # ignores any directory named vet

build: vet
	go build wiki.go
.PHONY:build # ignores any directory named build
