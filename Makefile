init:
	docker compose up -d

rebuild:
	docker comppose up -d --build

down:
	docker compose down