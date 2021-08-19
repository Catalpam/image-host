package main

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.PostForm("user")
		secret := ctx.PostForm("secret")
		if user !=  Config.User || secret != Config.Secret {
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
