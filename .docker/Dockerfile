FROM golang:1.21.0-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app/app/api

CMD ["air"]
