package main

import (
	"github.com/gin-gonic/gin"
)

func init()  {
	println()
}

func main() {
	const FolderLocation = "/"

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r = CollectRoute(r)
	panic(r.Run(":1111"))
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	//登录
	//r.POST("/minilogin", miniController.SetCookies)
	//路由组初始化
	imageRouter := r.Group("/")
	//使用Auth中间件进行权限认证
	imageRouter.Use(AuthMiddleWare())
	//小程序路由组
	imageRouter.POST("/upload", SaveImage)
	return r
}

