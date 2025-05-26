FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ../../go.mod go.sum ./
RUN go mod download

COPY ../.. .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/feed ./cmd/cli/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/feed .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

RUN apk add --no-cache postgresql-client

# Проверяем, что migrate доступен
RUN migrate -version

#ENTRYPOINT ["/app/feed"]