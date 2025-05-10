# Spar API Client

The Spar API Client is a Go library designed to interact with the [Spar Online](https://www.spar.si/online/) API.

[![go report card](https://goreportcard.com/badge/github.com/amadejkastelic/spar-api "go report card")](https://goreportcard.com/report/github.com/amadejkastelic/spar-api)
[![CI status](https://github.com/amadejkastelic/spar-api/actions/workflows/build.yaml/badge.svg?branch=main "test status")](https://github.com/amadejkastelic/spar-api/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/amadejkastelic/spar-api?tab=doc)

## Overview

* Supports listing products and categories.
* Supports pagination and filtering.
* Provides a command-line interface (CLI).


## Running

To run the provided example, use the following command:
```bash
nix run .#default -- --help
```

## Development

```bash
nix develop
```

## Build

```bash
nix build
```

## Test

To run tests, use the following commands:
```bash
nix develop
go test ./...
```

To run all checks, including tests, use:
```bash
nix flake check
```
