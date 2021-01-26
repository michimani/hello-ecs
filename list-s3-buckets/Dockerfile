FROM golang:1.15.5-alpine3.12 as build

ENV GOPATH=
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -o /main

ENTRYPOINT [ "/main" ]