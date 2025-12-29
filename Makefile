init: build \
	up

restart: down \
	init

up:
	docker compose up -d

build:
	docker compose build

down:
	docker compose down