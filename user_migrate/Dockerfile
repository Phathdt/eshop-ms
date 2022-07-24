FROM golang:1.18.3-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o migrate main.go

# Deployment environment
# ----------------------
FROM alpine:3.16
WORKDIR /app
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=builder --chown=nobody:nobody ./app/migrate /app/
COPY --from=builder --chown=nobody:nobody ./app/migrations /app/migrations/
COPY --from=builder --chown=nobody:nobody ./app/run.sh .

ENTRYPOINT sh run.sh
