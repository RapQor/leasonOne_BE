build:
	docker compose up --build

run:
	docker compose up

run-prod:
	docker compose up -d

down:
	docker compose down

main:
	go run main.go