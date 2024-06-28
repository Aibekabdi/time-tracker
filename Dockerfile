FROM golang:latest

WORKDIR /фзз

COPY go.* ./

RUN go mod download

COPY . ./


CMD [ "/cmd/app" ]