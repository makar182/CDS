FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o logging-service ./main.go

EXPOSE 8082
RUN chmod +x /app/logging-service
CMD ["/app/logging-service"]
