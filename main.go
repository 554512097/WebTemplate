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
	err := model.InitDB()
	if err != nil {
		log.Printf("database init failed, err: %v\n", err)
		return
	}
	e := gin.Default()
	params.InitCustomValidator()
	handler.InitRouter(*e)
	e.Run(":8080")
}
