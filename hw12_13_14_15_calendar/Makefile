BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"
POSTGRES_IMG="postgres:13.3"
POSTGRES_USER="pguser"
POSTGRES_PASSWORD="pgpass"
POSTGRES_DB="postgres"
DB_IMG_NAME="calendar-db"
MIGRATIONS_PATH="/db/migrations"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./configs/config.toml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

create-network:
	docker network create MyNetwork

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race -count=100 ./internal/...

lint:
	docker run --rm -v $(shell pwd):/app -w /app "golangci/golangci-lint:v1.49.0" golangci-lint run -v

db-up:
	docker run --hostname pgdb --network MyNetwork --name $(DB_IMG_NAME) -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -d $(POSTGRES_IMG)

db-down:
	docker kill $(DB_IMG_NAME) ; docker rm -f $(DB_IMG_NAME)

migrate:
	docker run --network MyNetwork --rm -v "$(shell pwd)/db/migrations":/flyway/sql flyway/flyway -url="jdbc:postgresql://pgdb:5432/postgres" -user=$(POSTGRES_USER) -password=$(POSTGRES_PASSWORD) -validateMigrationNaming="true" migrate

.PHONY: build run build-img  run-img version test lint
