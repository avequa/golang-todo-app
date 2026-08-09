# golang-todo-app

REST API на Go: стандартная библиотека, PostgreSQL, архитектура слои - структурированные логи с `request_id`, graceful shutdown, миграции и полное окружение в Docker Compose

## Стек

Go, PostgreSQL, `pgx/v5` с пулом соединений, `zap` для логов, `golang-migrate` для миграций, `envconfig` для конфигурации из окружения, Swagger для документации, `net/http` роутинг

## Архитектура

В `internal/core` инфраструктура, общая для всего приложения: логгер, пул соединений, HTTP-сервер, middleware
В `internal/features/users` - всё, что относится к пользователям
В `internal/features/tasks` - всё, что относится к задачам

```
cmd/app/main.go                — сборка зависимостей, единственное место, где слои встречаются
internal/core/                 — логгер, pgxpool, HTTP-сервер, middleware
internal/features/users/       — repository, service, transport
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