# PingPro

PingPro is a command-line network ping tool written in Go. It supports multiple protocols including TCP, UDP, and ICMP, allowing you to test connectivity to a host and port with configurable options.

## Features

- Supports TCP, UDP, and ICMP ping types
- Configurable timeout per ping
- Configurable interval between pings
- Verbose output option
- Optional colored output for better readability

## Installation

To install PingPro, you need to have Go installed on your system. Then you can build the project using:

```bash
go build -o pingpro main.go
```

Alternatively, you can get the package using:

```bash
go get github.com/csubbdev/pingpro
```

## Usage

```bash
./pingpro <host> <port> [flags]
```

### Flags

- `-t, --type` : Ping type to use. Options are `tcp`, `udp`, or `icmp`. Default is `tcp`.
- `-o, --timeout` : Timeout per ping in seconds. Default is 3 seconds.
- `-i, --interval` : Interval between pings in milliseconds. Default is 1000 ms.

### Example

Ping a host using TCP on port 80 with default settings:

```bash
./pingpro example.com 80
```

Ping a host using UDP on port 53 with a 5-second timeout output:

```bash
./pingpro example.com 53 -t udp -o 5
```

## License

This project is licensed under the MIT License.
