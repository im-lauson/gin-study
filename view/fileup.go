package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.POST("/upload", func(context *gin.Context) {
		//单文件
		file, _ := context.FormFile("file")
		fmt.Println(file.Filename)
		fmt.Println(file.Size / 1024)                       //单位是字节
		context.SaveUploadedFile(file, "./uploads/dog.jpg") //保存图片的路径
		context.JSON(200, gin.H{"msg": "上传成功"})
	})
	route.Run(":8080")
}
