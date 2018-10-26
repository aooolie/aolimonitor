package util

import (
	"net/http"
	"paging"
	"log"
	"fmt"
	"strings"
)

func HttpServer() {
	/* 设置访问的路由 */
	http.HandleFunc("/", paging.NodeInfo)
	http.HandleFunc("/onlinediscover", handleOnlineEvent)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleOnlineEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
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
	fmt.Println("the body is:", string(bodySlc))
	fmt.Fprintf(w, "fuck Wrold!") //这个写入到w的是输出到客户端的
}
