package handler

import (
	"log"
	"main/net"
	"main/params"

	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	userParams := params.UserParam{}
	err := ctx.ShouldBindJSON(&userParams)
	if err != nil {
		log.Printf("err: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, err.Error())
		return
	}
	if len(userParams.Nick) == 0 {
		userParams.Nick = "用户" + userParams.Nick
	}
	net.SucJson(ctx, userParams)
}

func login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	var err error
	err = defaultValidator().Var(account, "required,max=50")
	err = defaultValidator().Var(password, "required,max=30")
	if err != nil {
		log.Printf("err: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, "参数校验出错")
		return
	}
	net.SucJson(ctx, gin.H{
		"account":  account,
		"password": password,
	})
}

func updateUser(ctx *gin.Context) {
	userParams := params.UserParam{}
	err := ctx.ShouldBindJSON(&userParams)
	if err != nil {
		log.Printf("err: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, err.Error())
		return
	}
	net.SucJson(ctx, userParams)
}
