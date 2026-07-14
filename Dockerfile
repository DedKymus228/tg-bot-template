FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o tg-bot cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/tg-bot .
COPY .env .

CMD ["./tg-bot"]