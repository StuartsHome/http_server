## Introduction
HTTP server starts on port 8080.

The server accepts requests on the default route.

The port can be modified in the file: `main.go`

## How to run
To build and run a binary `http_server`:
```sh
make run
```

## Example request
```
> curl localhost:8080
User-Agent header needs to contain a value larger than 40 bytes, current value is of 10 bytes, curl/8.7.1
```

## How to test
```
make test
```
