FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download

RUN go install github.com/tsenart/vegeta@latest

COPY . .

RUN go build -o prime-backoff .

CMD ["./prime-backoff"]