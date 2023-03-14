package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.POST("/plus", func(context *gin.Context) {
		form, _ := context.MultipartForm() //多文件上传函数
		files, _ := form.File["upload[]"]  //File  map[string][]*FileHeader，上传的时候以此名为Key
		for _, file := range files {       //循环上传选中的文件
			context.SaveUploadedFile(file, "./uploads/"+file.Filename)
		}
		context.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传%d个文件", len(files))})
	})
	route.Run(":8080")
}
