package main

import (
	"fmt"
	"util"
	"config"
	)

func main() {
	rootPath, err := util.GetConfPath()
	var confPath string
	if err == nil {
		confPath = rootPath + config.MANAGE_PATH
	}
	ip := util.ReadConfContext(rootPath)
	fmt.Print(confPath, " ", ip)
}




