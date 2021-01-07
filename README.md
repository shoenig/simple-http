simple-http
===========

A toy http client and server implementation.

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/simple-http)](https://goreportcard.com/report/gophers.dev/cmds/simple-http)
[![GoDoc](https://godoc.org/gophers.dev/cmds/simple-http?status.svg)](https://godoc.org/gophers.dev/cmds/simple-http)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/simple-http.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/simple-http.svg)](LICENSE)

# Project Overview

Module `gophers.dev/cmds/simple-http` provides a trivial Go HTTP server
and client service implementation. By default the server lists for requests,
and the client makes requests.

# Getting Started

The `simple-http` package can be installed by running

```bash
$ go get gophers.dev/cmds/simple-http
```

# Example Usage

#### Running in server mode

The `simple-http` command can run in `server` mode

```bash
$ simple-http server
```

By default the server will bind to `127.0.0.1:8999`.

#### Running in client mode

The `simple-http` command can run in `client` mode

```bash
$ simple-http client
```

By default the client will make requests to `127.0.0.1:8999`.

# Configuration

#### Environment variables

The `simple-http` `server` and `client` modes can be configured through environment variables.

Use `BIND` and `PORT` to configure the `server` list address.

```bash
$ BIND=10.0.0.1 PORT=9000 simple-http server
```

Use `ADDRESS` and `PORT` to configure the `client` request address.

```bash
$ BIND=10.0.0.1 PORT=9000 simple-http client
```

#### CLI Flags

The `simple-http` `server` and `client` modes can be configured with command line flags.

Use `-bind` and `-port` to configure the `server` listen address.

```bash
$ simple-http server -bind 10.0.0.1 -port 9000
```

Use `-address` and `-port` to configure the `client` request address.

```bash
$ simple-http client -address 10.0.0.1 -port 9000
```

# Docker

#### Usage

Run as server with docker (listens at `0.0.0.0:8999`).

```bash
$ docker run --net=host --rm shoenig/simple-http:v1-amd64 server
```

Run as client with docker (requests to `127.0.0.1:8999`).

```bash
$ docker run --net=host --rm shoenig/simple-http:v1-amd64 client
```

#### Publish

Build image and push to docker hub.

`<version>` format is `v<n>-<arch>` (e.g. `v1-amd64`).

```
docker build -t shoenig/simple-http:<version> .
docker push shoenig/simple-http:<version>
```

# Contributing

The `gophers.dev/cmds/simple-http` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `gophers.dev/cmds/simple-http` module is open source under the [BSD-3-Clause](LICENSE) license.
