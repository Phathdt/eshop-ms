FROM golang:1.18.3-alpine
WORKDIR /usr/app

COPY go.mod .
COPY go.sum .
RUN go mod download

ENTRYPOINT go run main.go
