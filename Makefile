.PHONY: run dev-db-up

BINARY=app

test:
	go test

vet:
	go vet

migrate-db:
	cd migrations && goose postgres "user=${APP_DB_USER} password=${APP_DB_PASSWORD} host=${APP_DB_HOST}\
	 		port=${APP_DB_PORT}	dbname=postgres sslmode=disable" up

migrate-db-down:
	cd migrations && goose postgres "user=${APP_DB_USER} password=${APP_DB_PASSWORD} host=${APP_DB_HOST}\
	 		port=${APP_DB_PORT}	dbname=postgres sslmode=disable" down

generate:
	go generate ./...

vet:
	go vet ./...

migration:
	# name is name of migration passed as argument
	# make migration name=create_some_table
	cd migrations && goose create $(name) sql

build:
	go build -o target/${BINARY}

run:
	go run cmd/exchange/main.go

# development
server-up:
	docker compose up -d

server-down:
	cd dev-env && docker compose down

run-dev: dev-db-up run

