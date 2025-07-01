FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/. .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

RUN apk add --no-cache postgresql-client
RUN migrate -version