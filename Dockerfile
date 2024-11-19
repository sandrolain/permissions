# builder stage
FROM golang:1.23.3-alpine3.20 AS builder

RUN apk update && apk add --no-cache gcc git musl-dev

WORKDIR /app
COPY . .

RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod download &&\
  CGO_ENABLED=0 GOOS=linux go build -o /service ./cmd/main.go

# service stage
FROM alpine:3.20.3
COPY --from=builder /service /service
ENTRYPOINT ["/service"]
