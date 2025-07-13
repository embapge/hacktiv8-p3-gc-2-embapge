# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o payment-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/payment-service /app/payment-service
COPY .env .env

EXPOSE 8081
CMD ["./payment-service"]
