package main
import (
	"redis"
	"util"
	"time"
	"paging"
	"logger"
)

func main() {
	util.Start()
	redis.SetServerUrl2Redis()
	go paging.HttpServer()
	logger.Info("666")
	for {
		time.Sleep(100)
	}

}