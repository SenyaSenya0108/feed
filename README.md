# Микросервис фидов

## Деплой приложения

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