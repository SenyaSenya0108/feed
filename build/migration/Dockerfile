FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY migrations/. ./migrations

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

RUN migrate -version