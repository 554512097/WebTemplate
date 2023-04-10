package logger

import (
	"fmt"
	"log"
	"os"
)

var fileInput = true

func InitLogger() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	if fileInput {
		logFile, err := os.OpenFile("./logs/serve.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("open log file failed, err:", err)
			return
		}
		log.SetOutput(logFile)
	}
}
