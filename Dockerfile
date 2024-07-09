FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o kmip-health-checker

FROM alpine:3.16

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/kmip-health-checker .
COPY --from=builder /app/config ./config

EXPOSE 3322

CMD ["./kmip-health-checker", "serve"]
