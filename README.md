# kmip-health-checker

kmip-health checker is a go implementation of a simple healthchecker for a KMIP server.

It was built using the Flamingo web framework and the kmip-go library.

It creates a dynamic number of keys on a KMIP server and deletes them directly afterwards.

## Example

A fully configured ready-to-run example using pykmip as the KMIP server can be found
here: https://github.com/FriedrichRezner/kmip-health-checker-example

## Getting started

First step is having a running KMIP server accessible through your network.

I strongly suggest to try the ready-to-run example - but in case you want to set it up by hand, you can use the
following guide.

Second step is to set the following env variables:

| Env variable            | Description                                                                                         |
|-------------------------|-----------------------------------------------------------------------------------------------------|
| KMIP_SERVER_HOST        | Host of your KMIP server                                                                            |
| KMIP_SERVER_PORT        | Port of your KMIP server                                                                            |
| KMIP_SERVER_CIPHER_TYPE | Golang cipher type for the key and certificate (See https://go.dev/src/crypto/tls/cipher_suites.go) |
| KMIP_SERVER_CERT_FILE   | Path to the cert file                                                                               |
| KMIP_SERVER_KEY_FILE    | Path to the key file                                                                                |

Then you should execute

`go run main.go serve`

If everything worked, you can find a running KMIP health checker on your machine:

http://localhost:3322/health

## Unit Tests

You can run unit tests by executing

`go test ./...`