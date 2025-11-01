FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o hello-world

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/hello-world .
EXPOSE 8080
CMD ["./hello-world"]