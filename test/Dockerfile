FROM golang:1.11.1-stretch as builder
WORKDIR /test
COPY . /test
RUN rm -r .gotron .gotron-builder

ENV GO111MODULE=on

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./example -a example/main.go