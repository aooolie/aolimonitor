package paging

import (
	"net/http"
	"fmt"
)

func NodeInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	var bodySlc []byte = make([]byte, 1024)
	r.Body.Read(bodySlc)
	fmt.Println("the body is:", string(bodySlc))
	fmt.Fprintf(w, "Online Event Recived") //这个写入到w的是输出到客户端的
}