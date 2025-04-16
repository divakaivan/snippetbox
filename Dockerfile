FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o snippetbox ./cmd/web

# final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/snippetbox .
COPY ui ui

EXPOSE 4000

CMD ["./snippetbox", "-addr=:4000"]
