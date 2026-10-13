include .env
export

.PHONY: build run-http-app

run:
	go run main.go

build:
	docker build -t employees-app .

run-http-app:
	docker run --rm -p 5050:5050 employees-app