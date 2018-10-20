package main
import (
	"net/http"
	"log"
	"paging"
	"redis"
	"util"
)

func main() {
	util.Start()
	redis.SetServerUrl2Redis()

	http.HandleFunc("/", paging.NodeInfo) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}