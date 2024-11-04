CONTAINER = goauth
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

up:
	docker compose up -d
down:
	docker compose down --remove-orphans
down-v:
	docker compose down -v --remove-orphans
build:
	docker compose build --no-cache
exec:
	docker compose exec -it $(CONTAINER) bash
build-goose:
	docker compose exec -it $(CONTAINER) go build -o /app/goose-custom /app/cmd/migration/main.go
migrate: build-goose
	docker compose exec -it $(CONTAINER) ./goose-custom /app/migration up
migration: build-goose
	docker compose exec -it $(CONTAINER) goose create $(name) go -dir /app/migration
init: up migrate

create-test-db:
	docker compose exec -it postgres psql -U $(POSTGRES_USER) -c "DROP DATABASE IF EXISTS $(POSTGRES_TEST_DB);"
	docker compose exec -it postgres psql -U $(POSTGRES_USER) -c "CREATE DATABASE $(POSTGRES_TEST_DB);"

build-goose-test:
	docker compose exec -it $(CONTAINER) sh -c "GO_ENV="test" go build -o /app/goose-custom-test /app/cmd/migration/main.go"
test_migrate: create-test-db build-goose-test
	docker compose exec -it $(CONTAINER) sh -c "GO_ENV="test" ./goose-custom-test /app/migration up"
tests: test_migrate
	docker compose exec -it $(CONTAINER) sh -c "GO_ENV="test" go test ./..."