FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -ldflags="-s -w" -o /app/clone ./cmd/server/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/clone /app/clone

# Expose port 8080
EXPOSE 8080

CMD ["./clone"]
