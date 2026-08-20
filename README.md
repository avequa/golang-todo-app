# golang-todo-app

REST API на Go: `net/http`, PostgreSQL, `pgx` с пулом соединений, слоистая feature-based архитектура, структурированные логи с `request_id`, graceful shutdown, миграции, полное окружение в Docker Compose

## Стек

Go 1.26, `net/http`, PostgreSQL, `pgx/v5` + `pgxpool`, `zap` для структурированных логов, `golang-migrate` для миграций, `envconfig` для конфигурации из окружения, `validator` для валидации данных, Swagger/OpenAPI, Docker, Docker Compose, Makefile

## Архитектура

В `internal/core` — инфраструктура, общая для всего приложения

В `internal/features/users` - пользователи

В `internal/features/tasks` - задачи

В `internal/features/statistics` - статистика

каждый feature разделён на `repository`, `service` и `transport` слои

```text

cmd/todo-app/main.go          — сборка зависимостей и запуск приложения

internal/core/                — конфигурация, логгер, pgxpool, HTTP-сервер, middleware

internal/features/users/      — repository, service, transport
internal/features/tasks/      — repository, service, transport
internal/features/statistics/ — repository, service, transport

migrations/                   — миграции golang-migrate
docs/                         — Swagger/OpenAPI

```

`main.go` отвечает за dependency injection: создаёт repository, service и transport и связывает их между собой

## Логи и middleware

Каждый запрос проходит через цепочку `RequestID -> Logger -> Trace -> Panic`

На входе генерируется идентификатор запроса и кладётся в `context.Context` под типизированным ключом, дальше он попадает во все записи лога - по нему можно собрать всю историю одного запроса

Для логирования используется структурированный logger `zap`

## PostgreSQL и миграции

Для работы с PostgreSQL используется `pgx/v5` и пул соединений `pgxpool`

Миграции выполняются через `golang-migrate`:

```bash

make migrate-create seq=init
make migrate-up
make migrate-down

```

Данные PostgreSQL сохраняются в Docker volume

## Docker

Окружение запускается через Docker Compose

В compose используются отдельные контейнеры для приложения, PostgreSQL, миграций и port forwarding

Приложение собирается через multi-stage Dockerfile: Go используется только на этапе сборки, в итоговый Alpine-образ копируется готовый бинарный файл

## Запуск

```bash

git clone https://github.com/avequa/golang-todo-app.git

cd golang-todo-app

cp .env.example .env

make env-up          # запуск PostgreSQL
make migrate-up      # накатить миграции
make todo-app-run    # запуск приложения

```

Для запуска приложения и PostgreSQL в Docker:

```bash

make todo-app-deploy

```

API поднимется на `http://localhost:5050`

## API

Все эндпоинты живут под префиксом `/api/v1`, версия задаётся отдельным роутером

Для пользователей доступны создание (`POST /users`), список (`GET /users`), получение по идентификатору (`GET /users/{id}`), частичное обновление (`PATCH /users/{id}`) и удаление (`DELETE /users/{id}`)

Для задач доступны создание (`POST /tasks`), получение списка задач (`GET /tasks`), получение задачи по ID (`GET /tasks/{id}`), удаление задачи (`DELETE /tasks/{id}`) и частичное обновление (`PATCH /tasks/{id}`)

Для статистики доступен просмотр статистики (`GET /statistics`)
