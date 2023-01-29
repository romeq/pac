#general
GO=/bin/env go

#database
DB_USERNAME ?= pac
DB_PASSWORD ?= pac
DB_HOST ?= 127.0.0.1
DB_PORT ?= 5432
DB_NAME ?= pac
PG_STRING ?= postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

#paths
SERVER_PKG=./cmd/server

up: run

migrate:
	@-cat ./sqlc/models/* | psql -d $(PG_STRING) >/dev/null

migrate-down:
	@-psql -d $(PG_STRING) -f ./sqlc/down.sql >/dev/null

psql:
	@-psql -qd $(PG_STRING) 

clean-db: migrate-down migrate

run:
	@$(GO) run $(SERVER_PKG)

build:
	@$(GO) build -o server-binary $(SERVER_PKG)
