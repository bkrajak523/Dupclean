# Stage 1: Build
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o dupclean ./cmd/dupclean

# Stage 2: Run
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/dupclean /usr/local/bin/dupclean
ENTRYPOINT ["dupclean"]
