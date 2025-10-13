# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o hello-world

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/hello-world .
CMD ["./hello-world"]
