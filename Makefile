#!make
.PHONY: all test clean

SHELL := /bin/bash
NETWORK := $(shell docker network ls | grep profile-default)

compose-up:
ifeq ($(NETWORK),)
	docker network create profile-default
endif
	docker-compose up -d

migrate-create:
	migrate create -ext sql -dir ./internal/repository/migrations $(name)

test:
	go test ./... -race -count=1 -cover

run:
	go build -o bin/profile cmd/user.go
	./bin/profile

lint:
	golangci-lint run ./...