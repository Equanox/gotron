FROM golang:1.11.1-stretch as builder
WORKDIR /test
COPY example/ /test/example

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o example/main.go ./example