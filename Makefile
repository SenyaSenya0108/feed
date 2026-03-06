init: build \
	up

up:
	docker compose up -d

rebuild:
	docker compose up -d --build

restart: down \
	up

init:
	docker compose up -d

build:
	docker compose build

down:
	docker compose down

sync-data:
	docker compose run --rm cli sync-data