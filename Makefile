include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todo-app-postgres

env-down:
	@docker compose down todo-app-postgres

env-cleanup:
	@docker compose down todo-app-postgres
	@-rm -rf out/pgdata

env-port-forward-up:
	@docker compose up -d port-forwarder

env-port-forward-down:
	@docker compose down port-forwarder

migrate-create:

	@if [ -z "$(seq)" ]; then \
		echo "Нет параметра seq, пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \

	docker compose run --rm --user $(shell id -u):$(shell id -g) todo-app-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:

	@if [ -z "$(action)" ]; then \
		echo "Нет параметра action, пример: make migrate-action action=up"; \
		exit 1; \
	fi; \

	docker compose run --rm todo-app-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todo-app-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

todo-app-run:
	@go run cmd/todo-app/main.go

