#!make

SHELL := /bin/bash

test:
	go test ./... -race -count=1 -cover

lint:
	golangci-lint run ./...