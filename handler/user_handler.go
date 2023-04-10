package handler

import (
	"log"
	"main/net"
	"main/params"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(e gin.Engine) {
	e.POST("/user/register", register)
	e.POST("/user/login", login)
}

func register(ctx *gin.Context) {
	registerParams := params.RegisterParam{}
	err := ctx.ShouldBindJSON(&registerParams)
	if err != nil {
		log.Printf("err: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, err.Error())
		return
	}
	validateErr := params.Validator.Struct(registerParams)
	if validateErr != nil {
		log.Printf("validateErr: %v\n", validateErr)
		net.FailJson(ctx, net.CODE_FAILED, validateErr.Error())
		return
	}
	if len(registerParams.Nick) == 0 {
		registerParams.Nick = "用户" + registerParams.Nick
	}
	net.SucJson(ctx, registerParams)
}

func login(ctx *gin.Context) {
	up := new(params.UserParam)
	ctx.ShouldBindJSON(up)
	ctx.JSON(http.StatusOK, up)
}
