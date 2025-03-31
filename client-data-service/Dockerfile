FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETARCH
RUN echo "Building for architecture: $TARGETARCH"
RUN GOOS=linux GOARCH=arm64 go build -o /app/client-data-service ./cmd/main.go

# Минимальный финальный образ
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/client-data-service .

# Установим утилиту file
RUN apk add --no-cache file

# Проверка архитектуры бинарника
RUN file /root/client-data-service

# Убедитесь, что бинарник имеет права на выполнение
RUN chmod +x client-data-service

EXPOSE 8080

CMD ["./client-data-service"]
