# sz-sdk-go-mock

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage] where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: sz-sdk-go-mock is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `sz-sdk-go-mock` packages provide mock [Go]
objects representing the Software Development Kit that wraps the
Senzing C SDK APIs.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/sz-sdk-go-mock.svg)](https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-mock)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/sz-sdk-go-mock)](https://goreportcard.com/report/github.com/senzing-garage/sz-sdk-go-mock)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/LICENSE)

[![go-test-linux.yaml](https://github.com/senzing-garage/sz-sdk-go-mock/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/sz-sdk-go-mock/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/sz-sdk-go-mock/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/sz-sdk-go-mock/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/sz-sdk-go-mock/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/sz-sdk-go-mock/actions/workflows/go-test-windows.yaml)

## Overview

The Senzing `sz-sdk-go-mock` packages enable Go programs to simulate calling Senzing library functions.
The `sz-sdk-go-mock` implementation of the [sz-sdk-go] interface does not require the Senzing C libraries to be installed.

Other implementations of the [sz-sdk-go] interface include:

- [sz-sdk-go-core] - for calling Senzing Go SDK APIs natively
- [sz-sdk-go-grpc] - for calling Senzing SDK APIs over [gRPC]
- [go-sdk-abstract-factory] - An [abstract factory pattern] for switching among implementations

## Use

(TODO:)

## References

1. [Development]
1. [Errors]
1. [Examples]
1. [Package reference]

[abstract factory pattern]: https://en.wikipedia.org/wiki/Abstract_factory_pattern
[Development]: docs/development.md
[Errors]: docs/errors.md
[Examples]: docs/examples.md
[go-sdk-abstract-factory]: https://github.com/senzing-garage/go-sdk-abstract-factory
[Go]: https://go.dev/
[gRPC]: https://grpc.io/
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go-mock
[Senzing Garage]: https://github.com/senzing-garage-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[Senzing]: https://senzing.com/
[sz-sdk-go-core]: https://github.com/senzing-garage/sz-sdk-go-core
[sz-sdk-go-grpc]: https://github.com/senzing-garage/sz-sdk-go-grpc
[sz-sdk-go]: https://github.com/senzing-garage/sz-sdk-go
