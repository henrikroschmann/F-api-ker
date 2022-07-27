# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY src/. ./

RUN go build -o /f-api-ker

EXPOSE 8080

CMD [ "/f-api-ker" ]

