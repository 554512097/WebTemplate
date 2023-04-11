package net

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseVO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SucJson(ctx *gin.Context, data any) {
	bv := BaseVO{CODE_SUCCESS, "请求成功", data}
	ctx.JSON(http.StatusOK, bv)
}

func FailJson(ctx *gin.Context, code int, msg string) {
	bv := BaseVO{code, msg, nil}
	ctx.JSON(http.StatusOK, bv)
}
