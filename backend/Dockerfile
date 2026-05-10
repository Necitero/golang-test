# syntax=docker/dockerfile:1

# --- builder ---
FROM golang:1.25-bookworm AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# build single binary named "backend"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /usr/local/bin/backend .

# --- final ---
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app

# copy binary and make sure it's executable
COPY --from=builder /usr/local/bin/backend /usr/local/bin/backend
RUN chmod +x /usr/local/bin/backend

# create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8080
CMD ["/usr/local/bin/backend"]

