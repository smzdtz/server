package user

import (
	"smzdtz-server/pkg/app"
	"smzdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

var response = app.NewResponse()

// 用户列表
func Index(c *gin.Context) {
	response.Error(c, errcode.Success)
	// utils.Success(c, utils.Response{"data": true})
}

// 用户详情
func Detail(c *gin.Context) {
	response.Error(c, errcode.Success)
	// utils.Success(c, utils.Response{"data": true})
}

// 保存
func Store(c *gin.Context) {
	response.Error(c, errcode.Success)
	// utils.Success(c, utils.Response{"data": true})
}

// 验证用户登录
func ValidateLogin(c *gin.Context) {
	response.Error(c, errcode.Success)
	// utils.Success(c, utils.Response{"data": true})
}
