package System

import (
	"FTS/Error"
	"net"
)

func IPControl() string {

	addrs, err := net.InterfaceAddrs()

	Error.SystemExit(err)
	IP := ""
	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IP = ipnet.IP.String()
				break
			}

		}
	}
	return IP

}
