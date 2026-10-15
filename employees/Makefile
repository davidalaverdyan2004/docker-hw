include .env
export

.PHONY: build run-http-app

run:
	go run main.go

build:
	docker build -t employees-app .

run-http-app:
	docker run --rm -p 5050:5050 employees-app

postgres-up:
	@docker run -d --rm \
		--name employees-postgres \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-p 5432:5432 \
		-v employees-pgdata:/var/lib/postgresql \
		postgres:18

postgres-down:
	docker stop employees-postgres