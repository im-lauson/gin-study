package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMsg(err error, obj interface{}) string {
	//obj为结构体指针
	getObj := reflect.TypeOf(obj)
	//断言为具体类型，err是一个接口
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回，当找到一个第一个错误时，就可以结束了
			}
		}
	}
	return err.Error()
}

// 用户注册
type signUserInfo struct {
	Name       string `json:"name" binding:"required" msg:"Name是一个必填项哦！"` //用户名required：值不可为空，字段也不能少，某必填项
	Age        int    `json:"age" binding:"required" msg:"Age是一个必填项哦！"`
	Password   string `json:"password" msg:"密码是一个必填项哦！"`     //密码
	RePassword string `json:"rePassword" msg:"确认密码是一个必填项哦！"` //确认密码

}

func main() {
	route := gin.Default()
	route.POST("/", func(context *gin.Context) {

		var user signUserInfo
		//自带数据传输类型检验
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
			return
		}

		context.JSON(200, gin.H{"data": user})
	})
	route.Run(":8080")
}
