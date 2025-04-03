FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]