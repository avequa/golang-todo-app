# golang-todo-app

REST API на Go: `net/http`, PostgreSQL, `pgx` с пулом соединений, слоистая feature-based архитектура, структурированные логи с `request_id`, graceful shutdown, миграции, полное окружение в Docker Compose

## Стек

Go 1.26, `net/http`, PostgreSQL, `pgx/v5` + `pgxpool`, `zap` для структурированных логов, `golang-migrate` для миграций, `envconfig` для конфигурации из окружения, `validator` для валидации данных, Swagger/OpenAPI, Docker, Docker Compose, Makefile
Спека Swagger UI генерируется из аннотаций, после правки аннотаций нужен `make swagger-gen`

## Архитектура

В `/docs` - сваггер спека

В `internal/core` - инфраструктура, общая для всего приложения

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

Каждый запрос проходит через цепочку `CORS -> RequestID -> Logger -> Trace -> Panic`

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

## Переменные окружения

Конфигурация приложения задаётся через `.env` (создаётся из `.env.example`)

```dotenv
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s
ALLOWED_ORIGINS=http://localhost:5050,null

POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG

TIME_ZONE=UTC
```

`HTTP_ADDR` — адрес и порт, на котором слушает HTTP-сервер

`HTTP_SHUTDOWN_TIMEOUT` — таймаут graceful shutdown

`ALLOWED_ORIGINS` — список разрешённых origin для CORS

`POSTGRES_*` — параметры подключения к PostgreSQL

`LOGGER_LEVEL` — уровень логирования (`DEBUG`, `INFO`, и т.д.)

`TIME_ZONE` — таймзона приложения

## Локальный запуск

```bash

git clone https://github.com/avequa/golang-todo-app.git

cd golang-todo-app

cp .env.example .env

make swagger-gen      # генерация спеки
make env-up           # запуск PostgreSQL
make env-port-forward # пробросить порты (исп. для локальной разработки, когда часть окружения поднята в Docker)
make migrate-up       # накатить миграции
make todo-app-run     # запуск приложения 

```

## Для запуска приложения и PostgreSQL в Docker:

```bash

make todo-app-deploy

```

API поднимется на `http://localhost:5050`

## API

Все эндпоинты живут под префиксом `/api/v1`, версия задаётся отдельным роутером

Для пользователей доступны создание (`POST /users`), список (`GET /users`), получение по идентификатору (`GET /users/{id}`), частичное обновление (`PATCH /users/{id}`) и удаление (`DELETE /users/{id}`)

Для задач доступны создание (`POST /tasks`), получение списка задач (`GET /tasks`), получение задачи по ID (`GET /tasks/{id}`), удаление задачи (`DELETE /tasks/{id}`) и частичное обновление (`PATCH /tasks/{id}`)

Для статистики доступен просмотр статистики (`GET /statistics`)

## `todo-app-run` vs `todo-app-deploy`

Это два независимых способа запуска, которые не синхронизируются автоматически

`todo-app-run` - запускает приложение через `go run` напрямую с хоста
Статические файлы (например `public/index.html`) читаются с диска на каждый запрос, поэтому правки подхватываются сразу, без перезапуска

`todo-app-deploy` — собирает Docker-образ и запускает приложение в контейнере
Все файлы, включая `public/index.html`, копируются внутрь образа **на момент сборки**. После правок в коде или во фронтенде обязательно нужна пересборка:

```bash
make todo-app-deploy
```
