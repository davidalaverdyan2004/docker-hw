include .env
export

.PHONY: run build run-http-app postgres-up postgres-down

run:
	go run main.go

build:
	docker build -t employees-app .

run-http-app:
	docker compose up -d --build app

postgres-up:
	docker compose up -d postgres

postgres-down:
	docker compose down