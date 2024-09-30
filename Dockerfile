# STAGE 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o btree_service .

# STAGE 2: Run
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/btree_service .

EXPOSE 3000

CMD ["./btree_service"]
