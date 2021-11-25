package install

import (
	"smzdtz-server/internal/ecode"
	"smzdtz-server/internal/model"
	"smzdtz-server/internal/service"
	"smzdtz-server/pkg/app"
	"smzdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

var response = app.NewResponse()

type InstallForm struct {
	// DbType               string `json:"dbType" binding:"required,oneof=mysql postgres"`
	// DbHost               string `json:"dbHost" binding:"required,min=0,max=50"`
	// DbPort               int    `json:"dbPort" binding:"required,gt=0,lte=65535"`
	// DbUsername           string `json:"dbUsername" binding:"required,max=50"`
	// DbPassword           string `json:"dbPassword" binding:"required,max=30"`
	// DbName               string `json:"dbName" binding:"required,max=50"`
	// DbTablePrefix        string `json:"dbTablePrefix" binding:"max=20"`
	AdminUsername        string `json:"adminUsername" binding:"required,min=3"`
	AdminPassword        string `json:"adminPassword" binding:"required,min=6"`
	ConfirmAdminPassword string `json:"confirmAdminPassword" binding:"required,min=6"`
	AdminEmail           string `json:"adminEmail" binding:"required,email,max=50"`
}

// 安装状态
func Status(c *gin.Context) {
	response.Success(c, nil)
}

// 系统初始化
func Store(c *gin.Context) {
	var params InstallForm
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, errcode.ErrBind.WithDetails(err.Error()))
		return
	}
	if params.AdminPassword != params.ConfirmAdminPassword {
		response.Error(c, ecode.ErrTwicePasswordNotMatch)
		return
	}
	// 测试数据库连接
	// err := service.TestDbConnection(params)
	// if err != nil {
	// 	utils.Fail(c, utils.Response{"code": 10, "message": err})
	// 	return
	// }
	model.Db = service.CreateDb()
	// 创建数据库表
	migration := new(model.Migration)
	err := migration.Install()
	if err != nil {
		response.Error(c, errcode.ErrBind.WithDetails(err.Error()))
		return
	}
	// 创建管理员账号
	err = service.CreateAdminUser(params.AdminUsername, params.AdminPassword, params.AdminEmail)
	if err != nil {
		response.Error(c, errcode.ErrBind.WithDetails(err.Error()))
		return
	}

	response.Success(c, errcode.Success)
}
