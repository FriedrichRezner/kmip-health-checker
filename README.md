# KMIP Health Checker

KMIP Health Checker is a Go implementation of a simple health checker for a KMIP server.

This tool is built using the [Flamingo](https://github.com/i-love-flamingo) web framework and
the [kmip-go](https://github.com/ThalesGroup/kmip-go) library. It dynamically creates and then deletes a number of keys
on a KMIP server to verify that it is working properly.

A use case would be to deploy it in a Kubernetes cluster and setup a health check to ensure that a KMIP server inside
the cluster is functioning properly.

## Example

A fully configured, ready-to-run example repository using `pykmip` as the KMIP server can be
found [here](https://github.com/FriedrichRezner/kmip-health-checker-example).

# Prerequisites

The application requires Go version 1.22 or above.

## Getting Started

To get started, a KMIP server must be running and be accessible through the network.

The ready-to-run example repository is a fully functional implementation. But if manual setup is preferred, the
following steps have to be done.

### Configuration

Set the following environment variables:

| Environment Variable      | Description                                                                                                                 |
|---------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| `KMIP_SERVER_HOST`        | Host of the KMIP server                                                                                                     |
| `KMIP_SERVER_PORT`        | Port of the KMIP server                                                                                                     |
| `KMIP_SERVER_CIPHER_TYPE` | Golang cipher type for the key and certificate (See [Golang Cipher Suites](https://go.dev/src/crypto/tls/cipher_suites.go)) |
| `KMIP_SERVER_CERT_FILE`   | Path to the certificate file                                                                                                |
| `KMIP_SERVER_KEY_FILE`    | Path to the key file                                                                                                        |

### Running the Health Checker

After setting the environment variables, execute the following command:

```sh
go install
go run main.go serve
```

If you prefer using Docker, you can start the server using this command. You need also have to provide the correct
certificates.

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

If everything is configured correctly, the KMIP Health Checker can be accessed at:

[http://localhost:3322/health](http://localhost:3322/health)

## Unit Tests

To run unit tests, execute:

```sh
go test ./...
```
