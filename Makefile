#general
CGO_ENABLED ?= 1
GO=go

#database
DB_USERNAME ?= pac
DB_PASSWORD ?= pac
DB_HOST ?= 127.0.0.1
DB_PORT ?= 5432
DB_NAME ?= pac
PG_STRING ?= postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

#paths
SERVER_PKG=./cmd/server
SERVER_BIN=./server-bin

migrate:
	@-cat ./sqlc/models/* | psql -d $(PG_STRING) >/dev/null

migrate-down:
	@-psql -d $(PG_STRING) -f ./sqlc/down.sql >/dev/null

db-shell:
	@-psql -qd $(PG_STRING) 

db-reset: migrate-down migrate

build:
	@-CGO_ENABLED=$(CGO_ENABLED) $(GO) build -o $(SERVER_BIN) $(SERVER_PKG)

run: build
	@-$(SERVER_BIN)

clean:
	rm $(SERVER_BIN)
