# Spar API Client

## Overview

The Spar API Client is a Go library designed to interact with the [Spar Online](https://www.spar.si/online/) API. It provides a simple and efficient way to make requests to the API.

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
