# Simple Mock Server

This is a simple configurable mock server written in Go

## Usage

```bash
$ ./mock-server
$ ./mock-server -port=8081 -responseCode=500 -sleep=10
```

| Flag          | Optional | Default | Description                                        |
|---------------|----------|---------|----------------------------------------------------|
| -port         | Y        | 8080    | Port the server runs on                            |
| -responseCode | Y        | 200     | HTTP response code to return                       |
| -sleep        | Y        | 0       | Time to sleep (in seconds) before returning result |

## Building

Build for local use

```bash
$ go build
```

Build for linux use

```bash
$ env GOOS=linux GOARCH=amd64 go build
```
