#!make

SHELL := /bin/bash
NETWORK := $(shell docker network ls | grep profile-default)

compose-up:
ifeq ($(NETWORK),)
	docker network create profile-default
endif
	docker-compose up -d

test:
	go test ./... -race -count=1 -cover

run:
	go build -o profile ./...
	./profile

lint:
	golangci-lint run ./...