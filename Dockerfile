# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates mysql-client

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy migration files
COPY --from=builder /app/sql ./sql

# Create start script
RUN echo '#!/bin/sh' > start.sh && \
    echo 'echo "Waiting for database..."' >> start.sh && \
    echo 'while ! mysqladmin ping -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" --silent; do' >> start.sh && \
    echo '    sleep 1' >> start.sh && \
    echo 'done' >> start.sh && \
    echo 'echo "Database is ready!"' >> start.sh && \
    echo 'echo "Running migrations..."' >> start.sh && \
    echo 'mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < sql/migrations/001_create_orders_table.sql' >> start.sh && \
    echo 'echo "Migrations completed!"' >> start.sh && \
    echo 'echo "Starting application..."' >> start.sh && \
    echo 'exec ./main' >> start.sh && \
    chmod +x start.sh

EXPOSE 8000 8080 50051

CMD ["./start.sh"]
