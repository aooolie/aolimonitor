package util

import (
	"os"
	"os/exec"
	"bytes"
	"strings"
	"errors"
	"config"
	"io/ioutil"
)

var (
	CONFPATH string
	MASTERIP string
)

func GetConfPath() (string,error) {

	return unixHome()
}

func SetConfPath(path string) {
	CONFPATH = path
}

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

func ReadConfContext(path string) string {
	var confPath string
	confPath = path + config.MANAGE_PATH
	data, _ := ioutil.ReadFile(confPath)
	datas := string(data)
	subDatas := strings.Split(datas, "\n")
	for _, d := range subDatas {
		if strings.Contains(d, "#") {
			continue
		}
		arg := strings.Split(d, "=")
		if arg[0] == "master" {
			return strings.TrimSpace(arg[1])
		}
	}
	return ""
}
