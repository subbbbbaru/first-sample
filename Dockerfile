FROM golang:1.21-alpine AS builder

# Установка необходимых пакетов
# RUN apk add --no-cache git

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /go-webservice cmd/main.go

FROM alpine:latest

# RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /go-webservice .

EXPOSE 8080

CMD ["./go-webservice"]
