FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .