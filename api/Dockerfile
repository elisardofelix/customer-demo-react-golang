FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/.

FROM alpine:latest as production
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

FROM golang:1.20-alpine AS test
WORKDIR /app
COPY . .
RUN go mod download