FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN mkdir /logs

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/. .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cli ./cmd/cli/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/cli .

RUN apk add --no-cache postgresql-client

ENV TZ=Europe/Moscow

ENTRYPOINT ["./cli"]