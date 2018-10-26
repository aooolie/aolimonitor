package main
import (
	"redis"
	"util"
	"time"
)

func main() {
	util.Start()
	redis.SetServerUrl2Redis()
	go util.HttpServer()
	for {
		time.Sleep(100)
	}
}