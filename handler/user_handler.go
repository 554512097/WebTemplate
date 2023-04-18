package handler

import (
	"log"
	"main/model"
	"main/net"
	"main/utils"

	"github.com/gin-gonic/gin"
)

// @BasePath /user

// 注册
// @Summary 注册接口
// @Schemes
// @Description 用户注册
// @Tags register
// @Param 参数 body model.User true "用户模型"
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /user/register [post]
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
		log.Printf("校验错误: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, "参数校验出错")
		return
	}
	u := model.User{}
	err2 := model.GloableDB.Where("account = ?", account).First(&u).Error
	if err2 != nil {
		log.Printf("查询错误: %v\n", err2)
		net.FailJson(ctx, net.CODE_FAILED, "用户不存在")
		return
	}
	token, err3 := utils.GenerateJWT(u.ID)
	if err3 != nil {
		log.Printf("token错误: %v\n", err3)
		net.FailJson(ctx, net.CODE_FAILED, err3.Error())
		return
	}
	ctx.Header("token", token)
	net.SucJson(ctx, u)
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
		userParams.Nick = "用户" + userParams.Account
	}
	var dbError error
	if isModify {
		dbError = model.GloableDB.Update(&userParams).Error
	} else {
		u := model.User{}
		e := model.GloableDB.Where("account = ?", userParams.Account).First(&u).Error
		if e == nil {
			net.FailJson(ctx, net.CODE_FAILED, "用户已存在")
			return
		}
		dbError = model.GloableDB.Create(&userParams).Error
	}
	if dbError != nil {
		log.Printf("err: %v\n", dbError)
		net.FailJson(ctx, net.CODE_FAILED, dbError.Error())
		return
	}
	net.SucJson(ctx, userParams)
}

// Userinfo
// @Summary 获取用户信息接口
// @Schemes
// @Description 获取用户信息
// @Tags userinfo
// @Param id path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Router /user/info/{id} [get]
func userinfo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := defaultValidator().Var(id, "required,numeric")
	if err != nil {
		log.Printf("err: %v\n", err)
		net.FailJson(ctx, net.CODE_FAILED, "参数校验出错")
		return
	}
	var user model.User
	dbError := model.GloableDB.Where("id = ?", id).First(&user).Error
	if dbError != nil {
		log.Printf("err: %v\n", dbError)
		net.FailJson(ctx, net.CODE_FAILED, dbError.Error())
		return
	}
	net.SucJson(ctx, user)
}
