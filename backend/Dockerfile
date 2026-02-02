# build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# install ca-certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

# copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

# production stage
FROM scratch

WORKDIR /app

# copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy binary
COPY --from=builder /app/server .

# expose port
EXPOSE 8080

# run
ENTRYPOINT ["/app/server"]
