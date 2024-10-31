up:
	docker compose up -d
down:
	docker compose down --remove-orphans
down-v:
	docker compose down -v --remove-orphans
build:
	docker compose build --no-cache
exec:
	docker compose exec -it go-app bash
build-goose:
	docker compose exec -it go-app go build -o /app/goose-custom /app/cmd/migration/main.go
migrate: build-goose
	docker compose exec -it go-app ./goose-custom /app/migration up
migration: build-goose
	docker compose exec -it go-app goose create $(name) go -dir /app/migration
init: up migrate