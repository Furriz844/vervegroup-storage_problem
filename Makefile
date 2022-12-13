.PHONY: build
build:
	go build -o ./out/storage-api -v ./cmd/storage-api
.DEFAULT_GOAL := build