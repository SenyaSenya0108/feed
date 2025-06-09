# _Микросервис "Чат"_

## Деплой приложения

1. Заполнить .env

### Собрать и поднять приложение
```bash
  docker compose up -d
```

### Запустить в режиме отладки
```bash
  docker compose -f docker-compose.yaml -f docker-compose.debug.yaml up -d
```

### Опустить приложение
```bash
  docker compose down
```

### Создать миграцию
```bash
  docker compose run --rm migration sh -c 'migrate create -ext sql -dir migrations/ -seq create_hubs_table'
```

### Накатить миграции
```bash
  docker compose run --rm migration sh -c 'migrate -path migrations/ -database "$DB_URL" -verbose up'
```

### Откатить миграции

```bash
  docker compose run --rm migration sh -c 'migrate -path migrations/ -database "$DB_URL" -verbose down'
```

### Команда загрузки данных

```bash
  docker compose run --rm cli ./feed data-load
```