package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	route := gin.Default()
	route.POST("/co", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		readerFile, _ := file.Open()
		writerFile, _ := os.Create("./uploads/2.jpg")
		defer writerFile.Close()
		n, _ := io.Copy(writerFile, readerFile)
		fmt.Println(n)
		context.JSON(200, gin.H{"msg": "上传成功"})
	})
	route.Run(":8080")
}
