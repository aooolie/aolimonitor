package onlinediscover

import (
	"os"
	"os/exec"
	"bytes"
	"strings"
	"errors"
	"io/ioutil"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"util"
	"strconv"
)

var (
	CONFPATH string
	REDIS_HOST string
	REDIS_PORT string
)

func Start() {
	confPath, _ := getConfPath()
	ReadConfContext(confPath)
	getServerIP()

	/* 取第一个网卡ip地址作为主机地址 */
	util.HOST_IP = util.GetHostIP()[0]

	sendOnlineInfo()
	fmt.Print("Online event sent" + "\n")
}

func getConfPath() (string,error) {

	return unixHome()
}

func SetConfPath(path string) {
	CONFPATH = path
}

/* 获取config文件路径 */
func unixHome() (string,error) {
	//从环境变量中读取
	if home := os.Getenv("SOURCE"); home != "" {
		SetConfPath(home)
		return home, nil
	}
	//从命令行中读取
	var stdout bytes.Buffer
	cmd := exec.Command("sh","-c","eval echo ~$")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank path")
	}
	SetConfPath(result)
	return result, nil
}

/* 读取config文件配置信息 */
func ReadConfContext(path string) {
	var confPath string
	confPath = path + MANAGE_PATH
	data, _ := ioutil.ReadFile(confPath)
	datas := string(data)
	subDatas := strings.Split(datas, "\n")
	for _, d := range subDatas {
		if strings.Contains(d, "#") {
			continue
		}
		arg := strings.Split(d, "=")
		if strings.TrimSpace(arg[0]) == "REDIS HOST" {
			REDIS_HOST = strings.TrimSpace(arg[1])
			fmt.Print("REDIS_HOST: " + REDIS_HOST + "\n")
		}
		if strings.TrimSpace(arg[0]) == "REDIS PORT" {
			REDIS_PORT = strings.TrimSpace(arg[1])
			fmt.Print("REDIS_PORT: " + REDIS_PORT + "\n")
		}
	}
}

/* 从redis中读取server地址 */
func getServerIP() {

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

	Host = serverIP + ":" + serverPort

	defer c.Close()
}

func sendOnlineInfo() {
	var ans string
	parm := "host=" + util.HOST_IP + ",timestamp=" + strconv.FormatInt(time.Now().Unix(),10)
	for ans != "Online Event Recived" {
		ans = HttpPOST("/onlinediscover", parm)

		if ans != "Online Event Recived" {
			/* 如果请求不成功，过5s重新请求 */
			time.Sleep(5 * time.Second)
		}
	}
}