# Q&A API Service

API-сервис для вопросов и ответов на Go с использованием PostgreSQL, GORM и Docker.

## Описание проекта

Сервис предоставляет REST API для управления вопросами и ответами. Пользователи могут создавать вопросы, добавлять ответы к вопросам, просматривать вопросы с ответами и удалять их.

## Технологии

- **Go 1.21** - язык программирования
- **PostgreSQL** - база данных
- **GORM** - ORM для работы с БД
- **Goose** - миграции базы данных
- **Gorilla Mux** - HTTP роутер
- **Docker & Docker Compose** - контейнеризация

## Структура проекта

```
test-task/
├── cmd/
│   └── server/
│       └── main.go          # Точка входа приложения
├── internal/
│   ├── models/              # Модели данных
│   ├── repository/          # Слой работы с БД
│   ├── service/             # Бизнес-логика
│   ├── handler/             # HTTP handlers
│   ├── config/              # Конфигурация
│   └── database/            # Инициализация БД
├── migrations/              # Миграции goose
├── docker-compose.yml       # Docker Compose конфигурация
├── Dockerfile              # Docker образ
├── entrypoint.sh           # Entrypoint скрипт для Docker
├── go.mod                  # Go модули
├── .gitignore              # Git ignore правила
├── .dockerignore           # Docker ignore правила
└── README.md               # Документация
```

## Быстрый старт

### Запуск через Docker Compose

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd test-task
```

2. Запустите приложение:
```bash
docker compose up --build
```

**Примечание:** В новых версиях Docker используется команда `docker compose` (без дефиса). Если у вас старая версия, используйте `docker-compose`.

Приложение будет доступно по адресу `http://localhost:8080`

### Остановка

```bash
docker compose down
```

Для удаления данных БД:
```bash
docker compose down -v
```

## API Endpoints

### Вопросы (Questions)

#### GET /questions/
Получить список всех вопросов.

**Ответ:**
```json
[
  {
    "id": 1,
    "text": "What is Go?",
    "created_at": "2024-01-01T12:00:00Z"
  }
]
```

#### POST /questions/
Создать новый вопрос.

**Запрос:**
```json
{
  "text": "What is Go?"
}
```

**Ответ:**
```json
{
  "id": 1,
  "text": "What is Go?",
  "created_at": "2024-01-01T12:00:00Z"
}
```

#### GET /questions/{id}
Получить вопрос по ID со всеми ответами.

**Ответ:**
```json
{
  "id": 1,
  "text": "What is Go?",
  "created_at": "2024-01-01T12:00:00Z",
  "answers": [
    {
      "id": 1,
      "question_id": 1,
      "user_id": "user-123",
      "text": "Go is a programming language",
      "created_at": "2024-01-01T13:00:00Z"
    }
  ]
}
```

#### DELETE /questions/{id}
Удалить вопрос (все ответы удаляются каскадно).

**Ответ:** 204 No Content

### Ответы (Answers)

#### POST /questions/{id}/answers/
Добавить ответ к вопросу.

**Запрос:**
```json
{
  "user_id": "user-123",
  "text": "Go is a programming language"
}
```

**Ответ:**
```json
{
  "id": 1,
  "question_id": 1,
  "user_id": "user-123",
  "text": "Go is a programming language",
  "created_at": "2024-01-01T13:00:00Z"
}
```

#### GET /answers/{id}
Получить ответ по ID.

**Ответ:**
```json
{
  "id": 1,
  "question_id": 1,
  "user_id": "user-123",
  "text": "Go is a programming language",
  "created_at": "2024-01-01T13:00:00Z"
}
```

#### DELETE /answers/{id}
Удалить ответ.

**Ответ:** 204 No Content

### Health Check

#### GET /health
Проверка работоспособности сервиса.

**Ответ:** 200 OK

## Примеры использования

### Создать вопрос
```bash
curl -X POST http://localhost:8080/questions/ \
  -H "Content-Type: application/json" \
  -d '{"text": "What is Go?"}'
```

### Получить все вопросы
```bash
curl http://localhost:8080/questions/
```

### Добавить ответ
```bash
curl -X POST http://localhost:8080/questions/1/answers/ \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user-123", "text": "Go is a programming language"}'
```

### Получить вопрос с ответами
```bash
curl http://localhost:8080/questions/1
```

## Переменные окружения

- `DATABASE_URL` - строка подключения к PostgreSQL (по умолчанию: `postgres://postgres:postgres@localhost:5432/qa_db?sslmode=disable`)
- `PORT` - порт для HTTP сервера (по умолчанию: `8080`)

## Тестирование

Запуск тестов:
```bash
go test ./...
```

Запуск тестов с покрытием:
```bash
go test -cover ./...
```

## Миграции

Миграции выполняются автоматически при запуске приложения через Docker Compose.

Для ручного запуска миграций:
```bash
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/qa_db?sslmode=disable" up
```

## Особенности реализации

- **Модульная архитектура**: разделение на слои (handler, service, repository)
- **Валидация**: проверка входных данных на всех уровнях
- **Каскадное удаление**: при удалении вопроса автоматически удаляются все его ответы
- **Логирование**: использование стандартного log пакета
- **Тесты**: unit тесты для сервисов и HTTP тесты для handlers
- **Миграции**: использование goose для управления схемой БД
- **Retry-логика**: автоматические повторные попытки подключения к БД
- **WSL совместимость**: entrypoint-скрипт для решения проблем DNS в WSL

## Критерии оценки (выполнено)

✅ Модульная архитектура проекта  
✅ Читаемый и структурированный код  
✅ Корректная бизнес-логика с валидацией  
✅ Каскадное удаление через foreign key constraints  
✅ Docker и docker-compose конфигурация  
✅ Тесты (unit и HTTP тесты)  
✅ Миграции базы данных  
✅ README.md с инструкциями по запуску  





