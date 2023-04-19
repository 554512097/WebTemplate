package handler

import (
	"main/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(e gin.Engine) {
	//用户模块路由初始化
	initUserRouter(e)
}

func initUserRouter(e gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	ug := e.Group("/user")
	{
		ug.POST("/register", register)
		ug.POST("/login", login)
		ug.POST("/update", updateUser)
		ug.GET("/info", userinfo)
	}
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
