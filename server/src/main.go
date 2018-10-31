package main
import (
	"redis"
	"util"
	"time"
	"paging"
)

func main() {
	util.Start()
	redis.SetServerUrl2Redis()
	go paging.HttpServer()
	for {
		time.Sleep(100)
	}
}