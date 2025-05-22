FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY ../../go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY ../.. .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/feed ./cmd/cli/main.go

# Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/feed .

# Set the entrypoint
ENTRYPOINT ["/app/feed"]

# Default command
#CMD ["data-load"]