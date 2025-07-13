# Dockerfile - Auth Service (Go + MongoDB)
FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o shopping-service cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/shopping-service /app/shopping-service
COPY .env .env

EXPOSE 8081
CMD ["./shopping-service"]
