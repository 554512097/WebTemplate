package handler

import (
	"main/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(e gin.Engine) {
	//用户模块路由初始化
	initUserRouter(e)
}

func initUserRouter(e gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := e.Group("/api/v1")
	{
		ug := v1.Group("/user")
		{
			ug.POST("/register", register)
			ug.POST("/login", login)
			ug.POST("/update", updateUser)
		}
	}
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func defaultValidator() *validator.Validate {
	return binding.Validator.Engine().(*validator.Validate)
}
