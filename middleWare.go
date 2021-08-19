package main

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.PostForm("token")
		secret := ctx.PostForm("secret")

		if token != "hhh" || secret != "888" {
			ctx.JSON(413, gin.H{
				"code":  4001,
				"error": "Not Authorized!",
			})
			ctx.Abort()
			return
		}
		println("认证成功！")
	}
}
