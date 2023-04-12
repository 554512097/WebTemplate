package handler

import (
	"log"
	"main/model"
	"main/net"

	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) {
	modifyUser(ctx, false)
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
	modifyUser(ctx, true)
}

func modifyUser(ctx *gin.Context, isModify bool) {
	userParams := model.User{}
	err := ctx.ShouldBindJSON(&userParams)
	if err != nil {
		log.Printf("err: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, err.Error())
		return
	}
	if !isModify && len(userParams.Nick) == 0 {
		userParams.Nick = "用户" + userParams.Nick
	}
	var dbError error
	if isModify {
		dbError = model.GloableDB.Update().Error
	} else {
		dbError = model.GloableDB.Create(&userParams).Error
	}
	if dbError != nil {
		log.Printf("err: %v\n", dbError)
		net.FailJson(ctx, net.CODE_FAILED, dbError.Error())
		return
	}
	net.SucJson(ctx, userParams)
}
