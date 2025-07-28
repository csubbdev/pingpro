package ping

import (
	"fmt"
	"net"
	"time"
)

func TCP(host string, port int, timeout int) (string, bool) {
	address := FormatAddress(host, port)
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return fmt.Sprintf("\033[91mConnection to %s timed out", address), false
		}
		return fmt.Sprintf("\033[91m%s \033[0munreachable (%v)", address, err), false
	}
	defer conn.Close()
	return "", true
}
