# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app


COPY . .
RUN go mod download

# WORKDIR /app/src/server
RUN go build -o /build ./src/main.go 

EXPOSE 8080

CMD [ "/build" ]