package ping

import (
	"fmt"
	"time"

	gp "github.com/go-ping/ping"
)

func ICMP(host string, timeout int) (string, bool) {
	pinger, err := gp.NewPinger(host)
	if err != nil {
		return fmt.Sprintf("ICMP error: \033[91m%v\033[0m", err), false
	}
	pinger.Count = 1
	pinger.Timeout = time.Duration(timeout) * time.Second
	err = pinger.Run()
	if err != nil {
		return fmt.Sprintf("ICMP failed: \033[91m%v\033[0m", err), false
	}
	stats := pinger.Statistics()
	if stats.PacketsRecv < 1 {
		return fmt.Sprintf("ICMP no reply from \033[91m%s\033[0m", host), false
	}
	return "", true
}
