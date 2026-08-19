# golang-todo-app

REST API на Go: стандартная библиотека, PostgreSQL, архитектура слои - структурированные логи с `request_id`, graceful shutdown, миграции и полное окружение в Docker Compose

## Стек

Go, PostgreSQL, `pgx/v5` с пулом соединений, `zap` для логов, `golang-migrate` для миграций, `envconfig` для конфигурации из окружения, Swagger для документации, `net/http` роутинг

## Архитектура

В `internal/core` инфраструктура, общая для всего приложения
В `internal/features/users`      - всё, что относится к пользователям
В `internal/features/tasks`      - всё, что относится к задачам
В `internal/features/statistics` - всё, что относится к статистике

```
cmd/app/main.go                — сборка зависимостей, единственное место, где слои встречаются

internal/core/                 — логгер, доменные сущности, ошибки, pgxpool, HTTP-сервер, middleware

internal/features/users/       — repository, service, transport
internal/features/tasks/       — repository, service, transport
internal/features/statistics/  — repository, service, transport

migrations/                    — миграции golang-migrate
docs/                          — сгенерированная Swagger-спека
```

## Логи и middleware

Каждый запрос проходит через цепочку `RequestID -> Logger -> Trace -> Panic`
На входе генерируется идентификатор запроса и кладётся в `context.Context` под типизированным ключом, дальше он попадает во все записи лога - по нему можно собрать всю историю одного запроса

## Запуск

```bash
git clone https://github.com/avequa/golang-todo-app.git
cd golang-todo-app
cp .env.example .env

make env-up             # postgres + приложение
make migrate-up         # накатить миграции
make env-port-forward   # проброс портов
make todo-app-run       # запуск приложения
```

API поднимется на `http://localhost:8080`

## API

Все эндпоинты живут под префиксом `/api/v1`, версия задаётся отдельным роутером

Для пользователей доступны создание (`POST /users`), список (`GET /users`), получение по идентификатору (`GET /users/{id}`), частичное обновление (`PATCH /users/{id}`) и удаление (`DELETE /users/{id}`)

Для задач доступны создание (`POST /tasks`), получение списка задач (`GET /tasks`), получение задачи по ID (`GET /tasks/{id}`), удаление задачи (`DELETE /tasks/{id}`), частичное обновление задачи (`PATCH /tasks/{id}`)

Для статистики доступен просмотр статистики (`GET /statistics`)