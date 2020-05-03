.PHONY: build
build:
	go build -v ./cmd/notificator

.DEFAULT_GOAL := build
