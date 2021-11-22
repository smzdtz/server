package controllers

import (
	"fmt"
	"smzdtz-server/models"
	"smzdtz-server/services"
	"smzdtz-server/utils"

	"github.com/gin-gonic/gin"
)

// 安装状态
func Status(c *gin.Context) {
	utils.Success(c, utils.Response{"data": false})
}

// 系统初始化
func Store(c *gin.Context) {
	var params models.InstallForm
	if err := c.ShouldBindJSON(&params); err != nil {
		utils.Fail(c, utils.Response{"code": -1, "message": err.Error()})
		return
	}
	if params.AdminPassword != params.ConfirmAdminPassword {
		utils.Fail(c, utils.Response{"code": 10, "message": "两次输入密码不匹配"})
		return
	}
	// 测试数据库连接
	err := services.TestDbConnection(params)
	if err != nil {
		utils.Fail(c, utils.Response{"code": 10, "message": err})
		return
	}
	models.Db = services.CreateDb()
	// 创建数据库表
	migration := new(models.Migration)
	err = migration.Install()
	if err != nil {
		utils.Fail(c, utils.Response{"code": 10, "message": fmt.Sprintf("创建数据库表失败-%s", err.Error())})
		return
	}
	// 创建管理员账号
	err = createAdminUser(form)
	if err != nil {
		return json.CommonFailure("创建管理员账号失败", err)
	}

	utils.Success(c, utils.Response{"data": true})
}
