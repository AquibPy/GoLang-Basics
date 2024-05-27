# Website Health Checker

## Overview

**Website Health Checker** is a simple command-line tool written in Go that checks whether a specified website is running or down by attempting to establish a TCP connection to a given domain and port. 

## Features

- Checks the health status of a website.
- Customizable port number with a default to port 80.
- Displays whether the website is reachable or not, along with error details if unreachable.

## Installation

1. Ensure you have Go installed on your system. You can download and install Go from [golang.org](https://golang.org/).
2. Clone the repository or download the `check.go` and `main.go` files.
3. Build the application using the Go toolchain:

```sh
go build -o health-checker main.go check.go
```

## Usage

Run the built application from the command line, specifying the domain and optionally the port you want to check:

```sh
./health-checker --domain example.com [--port 80]
```

### Flags

- `--domain, -d`: The domain name to check (required).
- `--port, -p`: The port number to check (optional, defaults to 80).

### Examples

Check if `example.com` is reachable on the default port 80:

```sh
./health-checker --domain example.com
```

Check if `example.com` is reachable on port 8080:

```sh
./health-checker --domain example.com --port 8080
```

## Output

The tool will print the status of the specified domain:

- If the domain is reachable:
  ```
  [UP] example.com is reachable,
  From: <local address>
  TO: <remote address>
  ```
- If the domain is unreachable:
  ```
  [DOWN] example.com is unreachable,
  Error: <error details>
  ```

## Acknowledgments

- [urfave/cli](https://github.com/urfave/cli) - For the CLI framework used in this project.