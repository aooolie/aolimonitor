package onlinediscover

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/garyburd/redigo/redis"
)

var (
	Url string

)

func HttpPOST(url string, parm string) string{
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(parm))
	if err != nil {
		fmt.Print("error occur during create post request")
		return ""
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Print("error occur during post request")
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("error occur during reading body")
		return ""
	}
	return string(body)
}

/* 从redis中读取server地址 */
func setHttpServer() {

	/* 使用配置文件中的 Redis 服务器地址 */
	redisUrl := REDIS_HOST + ":" + REDIS_PORT

	c, err := redis.Dial("tcp", redisUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	serverIP, err := redis.String(c.Do("GET", SERVER_IP))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(serverIP)

	serverPort, err := redis.String(c.Do("GET", SERVER_PORT))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(serverPort)

	Url = serverPort + ":" + serverPort

	defer c.Close()
}