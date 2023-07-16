FROM golang:1.19-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
ENTRYPOINT go run ./cmd/server/main.go
