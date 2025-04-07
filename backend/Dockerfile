FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/joho/godotenv/cmd/godotenv@latest

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go mod tidy

RUN go build -modfile go.mod -o main main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]