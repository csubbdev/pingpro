# PingPro

PingPro is a command-line network ping tool written in Go. It supports multiple protocols including TCP, UDP, and ICMP, allowing you to test connectivity to a host and port with configurable options.

## Features

* Supports TCP, UDP, and ICMP ping types
* Configurable timeout per ping
* Configurable interval between pings

## Installation

To install PingPro, you need to have Go installed on your system. Then you can build the project using:

```bash
go build -o pingpro main.go
```

Alternatively, you can get the package using:

```bash
go get github.com/csubbdev/pingpro
```

## Adding PingPro to your system PATH

### Windows 11

1. Build the executable as shown above.
2. Move the `pingpro.exe` file to a directory of your choice, for example: `C:\Tools\pingpro\`
3. Add this directory to your system PATH environment variable:

   * Open the Start menu and search for "Environment Variables".
   * Click "Edit the system environment variables".
   * In the System Properties window, click "Environment Variables".
   * Under "System variables", find and select the "Path" variable, then click "Edit".
   * Click "New" and add the path to the directory where you placed `pingpro.exe` (e.g., `C:\Tools\pingpro\`).
   * Click OK on all dialogs to apply the changes.
4. Open a new Command Prompt window and type `pingpro` to run the tool from anywhere.

### Linux

1. Build the executable:

```bash
go build -o pingpro main.go
```

2. Move the `pingpro` binary to a directory in your PATH, for example:

```bash
sudo mv pingpro /usr/local/bin/
```

3. Make sure it's executable:

```bash
chmod +x /usr/local/bin/pingpro
```

4. Now you can use `pingpro` from any terminal.

## Building for Multiple Platforms

To build `pingpro` for multiple operating systems and architectures, install [`gox`](https://github.com/mitchellh/gox):

```bash
go install github.com/mitchellh/gox@latest
```

Then run:

```bash
gox -os="windows linux" -arch="amd64" -output="build/{{.OS}}_{{.Arch}}/pingpro"
```

This will create builds for Windows and Linux in the `build/` directory.

## Usage

```bash
pingpro <host> <port> [flags]
```

### Flags

* `-t, --type` : Ping type to use. Options are `tcp`, `udp`, or `icmp`. Default is `tcp`.
* `-o, --timeout` : Timeout per ping in seconds. Default is 3 seconds.
* `-i, --interval` : Interval between pings in milliseconds. Default is 1000 ms.

### Example

Ping a host using TCP on port 80 with default settings:

```bash
pingpro example.com 80
```

Ping a host using UDP on port 53 with a 5-second timeout output:

```bash
pingpro example.com 53 -t udp -o 5
```

## License

This project is licensed under the MIT License.
