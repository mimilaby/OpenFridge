# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY server/*.go ./server/

WORKDIR /app/server
RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]
