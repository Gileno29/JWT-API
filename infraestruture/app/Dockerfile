# Build stage
FROM golang:1.26 AS builder

WORKDIR /app

COPY ../../go.mod .
COPY ../../go.sum .

RUN go mod download

COPY ../../ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]