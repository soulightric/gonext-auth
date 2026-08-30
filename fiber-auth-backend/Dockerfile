# --- Stage 1: Build ---
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod & go.sum dulu biar layer cache lebih efisien
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Build binary statis, tanpa CGO biar image final bisa pakai base image minimal
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Stage 2: Run ---
FROM alpine:3.20

WORKDIR /app

# ca-certificates diperlukan kalau nanti backend manggil API eksternal (HTTPS)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]