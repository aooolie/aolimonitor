package util

import (
	"os"
	"os/exec"
	"bytes"
	"strings"
	"errors"
	"io/ioutil"
	"fmt"
	"logger"
)

var (
	CONFPATH string
	REDIS_HOST string
	REDIS_PORT string
)

func Start() {
	confPath, _ := getConfPath()
	ReadConfContext(confPath)
	logger.IS_DEBUG = true
	SetLogPath()
}

func getConfPath() (string,error) {
	confpath, _:= unixHome()
	fmt.Print("confPath: " + confpath + "\n")
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
	fmt.Print("confPath: " + confPath + "\n")
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

func SetLogPath() {
	os.Mkdir(CONFPATH + "/log", 0x666)
	logger.LOG_PATH = CONFPATH + "/log/monitor.log"
}