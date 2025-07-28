package ping

import (
	"fmt"
	"net"
)

func FormatAddress(host string, port int) string {
	if net.ParseIP(host) != nil && net.ParseIP(host).To4() == nil {
		return fmt.Sprintf("[%s]:%d", host, port) // IPv6
	}
	return fmt.Sprintf("%s:%d", host, port) // IPv4 or hostname
}
