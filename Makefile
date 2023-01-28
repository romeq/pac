GO=/bin/env go
DB_USERNAME ?= pac
DB_PASSWORD ?= pac
DB_HOST ?= 127.0.0.1
DB_PORT ?= 5432
DB_NAME ?= pac
PG_STRING ?= postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

migrate:
	@psql -d $(PG_STRING) -f ./sqlc/models/* >/dev/null

migrate-down:
	@psql -d $(PG_STRING) -f ./sqlc/down.sql >/dev/null

clean-db: migrate-down migrate

run:
	@$(GO) run ./cmd/account/main.go
