# Build stage
FROM golang:1.23.3 AS builder
WORKDIR /app
COPY go.mod .
COPY . .
RUN go build -o server .

# Run stage
FROM ubuntu:24.04
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8888
CMD [ "./server", "/tmp/server.txt", ":8888"]