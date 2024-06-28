FROM golang:latest

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./


CMD [ "/cmd/app" ]