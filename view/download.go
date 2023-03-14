package main

import "github.com/gin-gonic/gin"

func main() {

	route := gin.Default()
	route.GET("/getimg", func(context *gin.Context) {
		//context.File("./uploads/1.png") //有些时候不能唤起下载
		context.Header("Content-Type", "application/octet-stream")                     //表示是文件流，唤起浏览器下载，一般设置了这个就要设置文件名
		context.Header("Content-Disposition", "attachment; filename="+"gin-Study.png") //用来指定下载文件的文件名
		context.Header("Content-Transfer-Encoding", "binary")                          //表示传输过程中的编码形式，乱码问题可能就是因为它
		context.File("./uploads/1.png")                                                //单个有些时候不能唤起下载

	})

	route.Run(":8080")
}
