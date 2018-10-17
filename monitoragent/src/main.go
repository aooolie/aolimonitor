package main

import (
	"fmt"
	"onlinediscover"
)

func main() {
	rootPath, err := onlinediscover.GetConfPath()
	var confPath string
	if err == nil {
		confPath = rootPath + onlinediscover.MANAGE_PATH
	}
	ip := onlinediscover.ReadConfContext(rootPath)
	fmt.Print(confPath, " ", ip)
}
