#!make
.PHONY: all test clean

SHELL := /bin/bash
NETWORK := $(shell docker network ls | grep user-default)

compose-up:
ifeq ($(NETWORK),)
	docker network create user-networks
endif
	docker-compose up -d

migrate-create:
	migrate create -ext sql -dir ./internal/repository/migrations $(name)

test:
	go test ./... -race -count=1 -cover

run:
	go build -o bin/user cmd/*.go
	./bin/user server

lint:
	golangci-lint run ./...

docker-run:
	docker build --tag user .
	docker run --network=user_user-networks --env-file=test.env user server