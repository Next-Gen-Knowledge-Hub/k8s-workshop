# Build stage
FROM golang:1.23.3 AS builder
WORKDIR /app
COPY go.mod .
COPY . .
RUN go build -o arg-server .

# Run stage
FROM ubuntu:24.04
WORKDIR /app
COPY --from=builder /app/arg-server .
ENTRYPOINT ["./arg-server"]