FROM golang:1.24.5-alpine AS builder

RUN apk --no-cache add ca-certificates gcc g++ libc-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/app ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bin/app /bin/
COPY --from=builder /app/docker.env /.env

CMD ["/bin/app"]