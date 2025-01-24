# MPATH=app/databases/migrations
APP_NAME=auth-app
PG_USER=root
PG_PASSWORD=secret
PG_PORT=5433
DB_NAME=postgres
DB_URL=postgres://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  migrate-up        : Run all migrations"
	@echo "  migrate-down      : Rollback all migrations"
	@echo "  migrate-create    : Create a new migration file"
	@echo "  migrate-fix       : Fix dirty database version"
	@echo "Usage for migrate-create: make migrate-create name=<migration_name> dir=<directory>"
	@echo "  example: "
	@echo "  make migrate-create module=accounts name=create_users_table"
	@echo "  make migrate-up module=accounts step=1"
	@echo "  make migrate-down module=accounts step=1"
	@echo "  make migrate-fix module=accounts version=2"
	@echo "  make swag -> Generate Swagger docs"
	@echo "  make run -> Start backend server"

DSN = "$(DB_URL)&search_path=$(module)"

.PHONY: migrate-up
migrate-up:
	migrate -database $(DSN) -path modules/$(module)/migrations up $(step)

.PHONY: migrate-down
migrate-down:
	migrate -database $(DSN) -path modules/$(module)/migrations down $(step)

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir modules/$(module)/migrations $(name)
	
.PHONY: migrate-fix
migrate-fix:
	migrate -database $(DSN) -path modules/$(module)/migrations force $(version)

.PHONY: swag
swag:
	swag init --parseDependency -g ./main.go -o ./docs

# .PHONY: postgres-start
postgres-start:
	docker run --name $(APP_NAME) \
    -e POSTGRES_USER=$(PG_USER) \
    -e POSTGRES_PASSWORD=$(PG_PASSWORD) \
    -e POSTGRES_DB=$(DB_NAME) \
    -p $(PG_PORT):5432 \
    -d postgres:14

createdb:
	docker exec -it $(APP_NAME) createdb --username=$(PG_USER) --owner=$(PG_USER) $(DB_NAME)

dropdb:
	docker exec -it $(APP_NAME) dropdb $(DB_NAME)

postgres-stop:
	docker stop $(APP_NAME)

postgres-delete:
	docker rm $(APP_NAME)

.PHONY: run
run:
	nodemon --exec go run main.go --signal SIGTERM


# example:
# make migrate-create module=accounts name=create_users_table
# make migrate-up module=accounts step=1
# make migrate-down module=accounts step=1
# make migrate-fix module=accounts version=2
# make swag -> Generate Swagger docs
# make run -> Start backend server