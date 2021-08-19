package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
)

func SaveImage(ctx *gin.Context)  {

	f, err := ctx.FormFile("image")
	// 帖子标题
	title := ctx.PostForm("title")
	if title == "" {
		title = "Unkown"
	}
	// url前缀
	url := "https://image.fengzigeng.com/"
	// 文件夹目录
	fildDir := fmt.Sprintf(Config.Dir)

	if err != nil {
		ctx.JSON(400, gin.H{
			"code": 4001,
			"msg":  "上传失败!",
		})
		return
	} else {
		// 检测文件扩展名
		fileExt:=strings.ToLower(path.Ext(f.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
			ctx.JSON(400, gin.H{
				"code": 4002,
				"msg":  "Upload failed!Only PNG, JPG, GIF, JPEG files are allowed.",
			})
			return
		}

		// 生成文件名称： 文章题目+序号
		serial := 1
		fileName := fmt.Sprintf("%s-%d",title,serial)


		fmt.Printf("将要保存%s至%s\n",fileName,fildDir)

		//判断文件夹是否存在，若不存在则创建
		isExist,err:=IsFileExist(fildDir)
		if err != nil {
			ctx.JSON(400,gin.H{
				"code": 4004,
				"msg": "error:File Dir Read Error!",
			})
			panic("error:File Dir Read Error!")
		} else if !isExist{
			os.Mkdir(fildDir,os.ModePerm)
		}

		// 循环遍历以判断文件名是否已被占用
		for {
			if isExist,_ := IsFileExist(fmt.Sprintf("%s%s%s",fildDir,fileName,fileExt)); isExist == true {
				serial ++
				fileName = fmt.Sprintf("%s-%d",title,serial)
			} else {
				break
			}
		}

		filepath := fmt.Sprintf("%s%s%s",fildDir,fileName,fileExt)

		// 尝试保存文件，并抛出可能的错误
		errSave := ctx.SaveUploadedFile(f, filepath)
		if errSave != nil {
			ctx.JSON(400, gin.H{
				"code": 4005,
				"msg": "上传成功!",
			})
			panic("Save Image Error")
		}

		urlpath := url + fmt.Sprintf("%s-%s",fileName,fileExt)

		ctx.JSON(200, gin.H{
			"code": 200,
			"data":gin.H{
				"path":urlpath,
			},
			"msg": "上传成功!",
		})
	}
}

//func GetImage(c *gin.Context) {
//	imageName := c.Query("name")
//	imagePath := "./images/" + imageName
//	c.File(imagePath)
//}

//判断文件(夹)是否存在
func IsFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	// 文件大小为0也判断为不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}

	if err == nil {
		return true, nil
	}

	return false, err
}

