include .env
export $(shell sed 's/=.*//' .env)

MIGRATIONS_DIR=internal/migrations
MIGRATE=migrate -database "$(DATABASE_URL_LOCAL)" -path $(MIGRATIONS_DIR)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)
