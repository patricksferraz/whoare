.PHONY: build

build:
	docker-compose build whoare

.PHONY: status logs start stop clean

ps:
	docker-compose ps

logs:
	docker-compose logs -f whoare

up:
	docker-compose up -d

start:
	docker-compose start whoare

stop:
	docker-compose stop whoare

down:stop
	docker-compose down -v --remove-orphans

attach:
	docker-compose exec whoare bash

prune:
	docker system prune --all --volumes

.PHONY: gtest

gtest:
	go test -v -cover -coverprofile coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
