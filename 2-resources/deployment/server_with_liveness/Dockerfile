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
EXPOSE 9999
CMD [ "./server" ]