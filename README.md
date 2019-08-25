# Transavro Service

This is the Transavro service

Generated with

```
micro new github.com/nayanmakasare/transavro --namespace=cloudwalker --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: cloudwalker.srv.transavro
- Type: srv
- Alias: transavro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./transavro-srv
```

Build a docker image
```
make docker
```