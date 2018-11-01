package logger

import (
	"strings"
	"time"
	"os"
	"fmt"
)

var (
	IS_DEBUG bool
	LOG_PATH string
)

func Info(info string) {
	info = strings.Split(time.Now().String(), ".")[0] + " INFO: " + info + "\n"
	if IS_DEBUG {
		f, _ := os.OpenFile(LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x666)
		f.Write([]byte(info))
		f.Close()
	} else {
		fmt.Print(info)
	}
}

func Error(info string) {
	info = strings.Split(time.Now().String(), ".")[0] + " ERROR: " + info + "\n"
	f, _ := os.OpenFile(LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x666)
	f.Write([]byte(info))
	f.Close()
}

func Debug(info string) {
	if IS_DEBUG {
		info = strings.Split(time.Now().String(), ".")[0] + " DEBUG: " + info + "\n"
		f, _ := os.OpenFile(LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x666)
		f.Write([]byte(info))
		f.Close()
	}
}