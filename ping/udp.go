package ping

import (
	"fmt"
	"net"
	"time"
)

func UDP(host string, port int, timeout int) (string, bool) {
	address := FormatAddress(host, port)
	conn, err := net.DialTimeout("udp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return fmt.Sprintf("\033[91mConnection to %s timed out", address), false
		}
		return fmt.Sprintf("\033[91m%s \033[0munreachable (%v)", address, err), false
	}
	defer conn.Close()

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return fmt.Sprintf("UDP send error: \033[91m%v\033[0m", err), false
	}

	conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return fmt.Sprintf("No UDP response from \033[91m%s\033[0m", address), false
	}

	return "", true
}
