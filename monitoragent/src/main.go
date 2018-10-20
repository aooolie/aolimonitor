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
	fmt.Print(confPath, " ", ip)

	/* 上报上线事件 */

}
