package main

import (
	"log"
	"main/handler"
	"main/logger"
	"main/model"
	"main/params"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	logger.InitLogger()
	err := model.InitDB()
	if err != nil {
		log.Printf("database init failed, err: %v\n", err)
		return
	}
	e := gin.Default()
	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	params.InitCustomValidator()
	handler.InitRouter(*e)
	e.Run(":8080")
}
