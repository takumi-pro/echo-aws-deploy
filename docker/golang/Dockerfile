# =========== build stage ===========
FROM golang:1.19.12-alpine3.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/api/main.go

# =========== production stage =============
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
CMD ["/app/main"]