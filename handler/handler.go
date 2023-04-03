package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "this is the index page")
}

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "hello go",
	})
}
