# KMIP Health Checker 
[![Go Reference](https://pkg.go.dev/badge/github.com/friedrichrezner/kmip-health-checker.svg)](https://pkg.go.dev/github.com/friedrichrezner/kmip-health-checker)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/friedrichrezner/kmip-health-checker)
[![Go Report Card](https://goreportcard.com/badge/github.com/friedrichrezner/kmip-health-checker)](https://goreportcard.com/report/github.com/friedrichrezner/kmip-health-checker)

KMIP Health Checker is a Go implementation of a simple health checker for a KMIP server.

This tool is built using the [Flamingo](https://github.com/i-love-flamingo/flamingo) web framework and
the [kmip-go](https://github.com/ThalesGroup/kmip-go) library.

It dynamically creates and then deletes a number of keys on a KMIP server to verify that it is working properly. A use
case would be to deploy it in an environment with a KMIP server to setup an external health check to ensure that a
KMIP server is functioning properly.

## Example

A fully configured, ready-to-run example repository using `pykmip` as the KMIP server can be
found [here](https://github.com/FriedrichRezner/kmip-health-checker-example).

## Getting Started

The application requires at least Go version 1.22 or a Docker environment. A KMIP server must be running and
accessible through the network.

The [ready-to-run example repository](https://github.com/FriedrichRezner/kmip-health-checker-example) is a fully
functional implementation. For manual setup, the
following steps have to be done.

Set the following environment variables:

| Environment Variable      | Description                                                                                                                 |
|---------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| `KMIP_SERVER_HOST`        | Host of the KMIP server                                                                                                     |
| `KMIP_SERVER_PORT`        | Port of the KMIP server                                                                                                     |
| `KMIP_SERVER_CIPHER_TYPE` | Golang cipher type for the key and certificate (See [Golang Cipher Suites](https://go.dev/src/crypto/tls/cipher_suites.go)) |
| `KMIP_SERVER_CERT_FILE`   | Path to the certificate file                                                                                                |
| `KMIP_SERVER_KEY_FILE`    | Path to the key file                                                                                                        |

Then you can either run it locally or start it in a Docker container.

### Locally

After setting the environment variables above, execute the following command:

```sh
go mod download
go build -o kmip-health-checker
./kmip-health-checker serve
```

If everything is configured correctly, the KMIP Health Checker can be accessed
at: [http://localhost:3322/health](http://localhost:3322/health)

### Docker

```sh
docker build -t kmip-health-checker .
docker run -p 3322:3322 \
-e KMIP_SERVER_HOST=your_kmip_server_host \
-e KMIP_SERVER_PORT=your_kmip_server_port \
-e KMIP_SERVER_CIPHER_TYPE=123 \
-e KMIP_SERVER_CERT_FILE=your_cert_file_path \
-e KMIP_SERVER_KEY_FILE=your_key_file_path \
kmip-health-checker
```

If everything is configured correctly, the KMIP Health Checker can be accessed
at: [http://localhost:3322/health](http://localhost:3322/health)

## Unit Tests

To run unit tests, execute:

```sh
go test ./...
```

When interfaces are changed during development, mocks have to be regenerated. To do this, execute
the [doc.go](doc.go) file in the directory:

```sh
go generate doc.go
```

## OpenAPI documentation

A complete OpenAPI specification with available requests, parameters and responses can be found in the root of the
directory: [openapi.yaml](openapi.yaml).
