package main

import (
	"main/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.Index)
	r.GET("/hello", handler.Hello)
	r.Run(":8080")
}
