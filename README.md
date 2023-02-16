# g2-sdk-go-mock

## :warning: WARNING: g2-sdk-go-mock is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing g2-sdk-go-mock packages provide mock
[Go](https://go.dev/)
objects representing the Software Development Kit that wraps the
Senzing C SDK APIs.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/g2-sdk-go-mock.svg)](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/g2-sdk-go-mock)](https://goreportcard.com/report/github.com/senzing/g2-sdk-go-mock)
[![go-test.yaml](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test.yaml)

## Overview

The Senzing `g2-sdk-go-mock` packages enable Go programs to simulate calling Senzing library functions.
The `g2-sdk-go-mock` implementation of the
[g2-sdk-go](https://github.com/Senzing/g2-sdk-go)
interface.

Other implementations of the
[g2-sdk-go](https://github.com/Senzing/g2-sdk-go)
interface include:

- [g2-sdk-go-base](https://github.com/Senzing/g2-sdk-go-base) - for
  calling Senzing Go SDK APIs natively
- [g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc) - for
  calling Senzing SDK APIs over [gRPC](https://grpc.io/)
- [go-sdk-abstract-factory](https://github.com/Senzing/go-sdk-abstract-factory) - An
  [abstract factory pattern](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
  for switching among implementations

## Use

(TODO:)

## Development

### Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

### Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=g2-sdk-go-mock
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

### Test

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean test

## Error prefixes

Error identifiers are in the format `senzing-PPPPnnnn` where:

`P` is a prefix used to identify the package.
`n` is a location within the package.

Prefixes:

1. `6031` - g2config
1. `6032` - g2configmgr
1. `6033` - g2diagnostic
1. `6034` - g2engine
1. `6035` - g2hasher
1. `6036` - g2product
1. `6037` - g2ssadm
