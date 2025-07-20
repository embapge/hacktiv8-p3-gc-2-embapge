# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

# Copy root module files and proto for replace resolution
WORKDIR /app
COPY go.mod go.sum ./
COPY proto ./proto

# Copy service module files and source, then build
WORKDIR /app/payment-service
COPY payment-service/go.mod payment-service/go.sum ./
COPY payment-service/. ./
# Download dependencies and build
RUN go mod download && \
    go build -o payment-service cmd/main.go

FROM alpine:latest
WORKDIR /app
# Copy built binary
COPY --from=build /app/payment-service/payment-service ./payment-service
# Copy service env
COPY payment-service/.env .env

EXPOSE 8081
CMD ["./payment-service"]
