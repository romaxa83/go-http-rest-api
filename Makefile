.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: db_into
db_into:
	docker-compose exec db sh

.PHONY: up
up: up-container permission

.PHONY: up-container
up-container:
	docker-compose up -d

.PHONY: permission
permission:
	sudo chmod 777 -R docker/db

.PHONY: down
down:
	docker-compose down --remove-orphans

.DEFAULT_GOAL := build