package main

import (
	"log"
	"main/handler"
	"main/logger"
	"main/model"
	"main/params"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	b, err := model.InitDataBase()
	if err != nil || !b {
		log.Printf("err: %v\n", err)
		log.Println("database init failed")
	}
	e := gin.Default()
	params.InitCustomValidator()
	handler.InitRouter(*e)
	e.Run(":8080")
}
