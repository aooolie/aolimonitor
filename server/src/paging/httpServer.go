package paging

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"redis"
)

func HttpServer() {
	/* 设置访问的路由 */
	http.HandleFunc("/", NodeInfo)
	http.HandleFunc("/onlinediscover", handleOnlineEvent)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//TODO 这里HandleFunc是不是并发的
func handleOnlineEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	/* 解析URL传递的参数 */
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	var bodySlc []byte = make([]byte, 1024)

	r.Body.Read(bodySlc)
	parms := strings.Split(string(bodySlc[:]), ",")

	var timestamp []byte
	var agentIP string
	for _, parm := range parms {

		body := strings.Split(parm,"=")

		if body[0] == "timestamp" {
			//TODO 为什么这里body长度是994
			timestamp = []byte(body[1])[0:9]
		}
		if body[0] == "host" {
			agentIP = body[1]
		}
	}
	/* 将agentIP写入Redis */
	redis.SetAgentIP2Redis(agentIP, timestamp)
	fmt.Println("the body is:", string(bodySlc))
	fmt.Fprintf(w, "Online Event Recived") //这个写入到w的是输出到客户端的
}
