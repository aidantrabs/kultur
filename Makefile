.PHONY: dev db-up db-down migrate migrate-down migrate-status build

# database
db-up:
	docker compose up -d

db-down:
	docker compose down

# backend
dev:
	cd backend && go run ./cmd/server

# migrations
migrate:
	goose -dir backend/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir backend/migrations postgres "$(DATABASE_URL)" down

migrate-status:
	goose -dir backend/migrations postgres "$(DATABASE_URL)" status

# build
build:
	cd backend && go build -o bin/server ./cmd/server

# sqlc
sqlc:
	cd backend && sqlc generate
