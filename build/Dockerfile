FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY .env ./
COPY . .
RUN go build -o dc-watch-service ./cmd/main.go
CMD ["./dc-watch-service"]
