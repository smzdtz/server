package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// json 格式输出

type Response map[string]interface{}

// 0 正确返回
// > 0 调用OpenAPI时发生错误，需要开发者进行相应的处理
// -1 请求参数无效
// -2 IP没有权限

// 成功时返回
func Success(ctx *gin.Context, res Response) {

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    res["data"],
	})
}

// 函数名称, -1请求参数无效,-2IP没有权限,>1业务错误
func Fail(ctx *gin.Context, res Response) {
	code := -1
	if val, err := res["code"]; err {
		code = val.(int)
	}
	message := "fail"
	if val, err := res["message"]; err {
		message = val.(string)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
