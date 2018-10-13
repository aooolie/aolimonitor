package util

import (
	"os"
	"os/exec"
	"bytes"
	"strings"
	"errors"
)

func GetConfigPath() (string,error) {

	return unixHome()
}

func unixHome() (string,error) {
	//从环境变量中读取
	if home := os.Getenv("GOPATH"); home != "" {
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
	return result, nil
}