.PHONY: help db-up db-down db-logs db-shell dev-up dev-down build run

# Available commands help
help:
	@echo "Available commands:"
	@echo "  build         - Build the telegram-bot"
	@echo "  run           - Run the telegram-bot"
	@echo "  run-with-db   - Run the telegram-bot with PostgreSQL"
	@echo "  db-up         - Start PostgreSQL"
	@echo "  db-down       - Stop PostgreSQL"
	@echo "  dev-up        - Start PostgreSQL + pgAdmin"
	@echo "  dev-down      - Stop all development services"

# Build and run application
build:
	go build -v -o bin/telegram-bot.exe ./cmd/bot

run:
	go run ./cmd/telegrambot/main.go

run-with-db: db-up
	go run ./cmd/telegrambot/main.go

# Database management
db-up:
	docker-compose up -d postgres

db-down:
	docker-compose down

# Development with pgAdmin
dev-up:
	docker-compose --profile dev up -d

dev-down:
	docker-compose --profile dev down
