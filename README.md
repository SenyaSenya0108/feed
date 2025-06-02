# Микросервис фидов

## Деплой приложения

1. Заполнить .env

### Собрать и поднять приложение
```bash
  docker compose up -d
```

### Опустить приложение
```bash
  docker compose down
```

### Накатить миграции
```bash
  docker compose run --rm cli sh -c 'migrate -path migrations/ -database "$DB_URL" -verbose up'
```

### Откатить миграции 

```bash
  docker compose run --rm cli sh -c 'migrate -path migrations/ -database "$DB_URL" -verbose down'
```

### Команда загрузки данных

```bash
  docker compose run --rm cli ./feed data-load
```