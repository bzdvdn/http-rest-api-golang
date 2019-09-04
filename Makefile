.PHONY: build
build:
	go build -v ./cmd/apiserver

.DEAFAULT_GOAL := build