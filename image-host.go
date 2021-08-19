package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func init()  {
	gin.SetMode(gin.ReleaseMode)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("读取config.yaml出错: %v", err)
	}
	if yaml.Unmarshal(yamlFile, &Config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	fmt.Printf("Init Config:\n")
	fmt.Printf("-----User: %s\n",Config.User)
	fmt.Printf("-----Secret: %s\n",Config.Secret)
	fmt.Printf("-----Dir: %s\n\n",Config.Dir)
	fmt.Println("Slim Config:")
	fmt.Printf("-----Enable: %s\n",Config.Slim.IsEnable)
	fmt.Printf("-----Threshold: %s\n\n",Config.Slim.Pixels)
	fmt.Println("Word Watermark Config:")
	fmt.Printf("-----Enable: %s\n",Config.Watermark.Word.IsEnable)
	fmt.Printf("-----Content: %s\n\n",Config.Watermark.Word.Content)
	fmt.Println("Image Watermark Config:")
	fmt.Printf("-----Enable: %s\n",Config.Watermark.Image.IsEnable)
	fmt.Printf("-----ImagePath: %s\n\n\n",Config.Watermark.Image.Path)
}

func main() {
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

