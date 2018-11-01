package util

import (
	"net"
	"fmt"
	"os"
)

/* 获取 Server 地址 */
func GetHostIP() []string {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		fmt.Println("cant get host of this node, exit 1")
		os.Exit(1)
	}
	for _, address:= range addrs {
		if ipnet,ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return ips
}


