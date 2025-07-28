package main

import (
	"fmt"
	"os"
	"pingga/ping"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var (
		pingType string
		timeout  int
		interval int
	)

	pflag.StringVarP(&pingType, "type", "t", "tcp", "Ping type: tcp, udp, icmp")
	pflag.IntVarP(&timeout, "timeout", "o", 3, "Timeout per ping (sec)")
	pflag.IntVarP(&interval, "interval", "i", 1000, "Interval (ms)")

	pflag.Parse()

	args := pflag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: ./pingpro <host> <port> [-t tcp|udp|icmp] [-o timeout] [-i interval]")
		os.Exit(1)
	}

	host := args[0]
	portStr := args[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port number.")
		os.Exit(1)
	}

	fmt.Print("\033[2J\033[H")

	for {
		start := time.Now()
		var result string
		success := false

		switch strings.ToLower(pingType) {
		case "tcp":
			result, success = ping.TCP(host, port, timeout)
		case "udp":
			result, success = ping.UDP(host, port, timeout)
		case "icmp":
			result, success = ping.ICMP(host, timeout)
		default:
			fmt.Println("Invalid protocol:", pingType)
			os.Exit(1)
		}

		elapsed := float64(time.Since(start).Milliseconds())

		if success {
			line := fmt.Sprintf("\033[0mConnected to \033[92m%s\033[0m time=\033[92m%.2f ms\033[0m protocol=\033[92m%s\033[0m port=\033[92m%d\033[0m", host, elapsed, strings.ToUpper(pingType), port)
			fmt.Printf("%s\n", line)
		} else {
			fmt.Printf("%s\n", result)
		}

		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
