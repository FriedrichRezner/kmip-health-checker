# Basis-Image mit Golang und Alpine Linux
FROM golang:1.22.5-alpine as builder

# Setzen des Arbeitsverzeichnisses
WORKDIR /app

# Kopieren der Go-Moduldateien und Abhängigkeiten installieren
COPY go.mod go.sum ./
RUN go mod download

# Kopieren des gesamten Quellcodes
COPY . .

# Builden der Flamingo-Anwendung
RUN go build -o kmip-health-checker

# Neues, minimales Image für die Ausführung der Anwendung
FROM alpine:3.16

# Installieren der notwendigen CA-Zertifikate
RUN apk add --no-cache ca-certificates

# Setzen des Arbeitsverzeichnisses
WORKDIR /app

# Kopieren des gebauten Binaries aus dem vorherigen Schritt
COPY --from=builder /app/kmip-health-checker .
COPY --from=builder /app/config ./config

# Exponieren des Ports auf dem der Dienst läuft (passen Sie dies bei Bedarf an)
EXPOSE 3322

# Starten der Anwendung mit dem Befehl "serve"
CMD ["./kmip-health-checker", "serve"]
