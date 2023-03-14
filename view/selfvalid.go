package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func _GetValidMsg(err error, obj interface{}) string {
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
type _signUserInfo struct {
	Name string `json:"name" binding:"required,sign" msg:"Name是一个必填项哦！"` //用户名required：值不可为空，字段也不能少，某必填项
	Age  int    `json:"age" binding:"required" msg:"Age是一个必填项哦！"`
}

// 自定义一个绑定器
func singValid(fl validator.FieldLevel) bool {
	//包含字段
	var nameList []string = []string{"didi", "dada", "xixi"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}

func main() {
	route := gin.Default()
	//自定义一个Tag绑定器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", singValid)
	}

	route.POST("/", func(context *gin.Context) {

		var user _signUserInfo
		//自带数据传输类型检验
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}

		context.JSON(200, gin.H{"data": user})
	})
	route.Run(":8080")
}
