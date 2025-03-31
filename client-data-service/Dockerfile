FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o client-data-service ./main.go

EXPOSE 8080
RUN chmod +x /app/client-data-service
CMD ["/app/client-data-service"]