package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitRouter(e gin.Engine) {
	//用户模块路由初始化
	initUserRouter(e)
}

func initUserRouter(e gin.Engine) {
	userGroup := e.Group("/user")
	userGroup.POST("/register", register)
	userGroup.POST("/login", login)
	userGroup.POST("/update", updateUser)
}

func defaultValidator() *validator.Validate {
	return binding.Validator.Engine().(*validator.Validate)
}
