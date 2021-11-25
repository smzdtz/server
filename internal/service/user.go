package service

import (
	// "smzdtz-server/models"
	// "smzdtz-server/utils"

	"smzdtz-server/internal/model"
	// "smzdtz-server/pkg/string"
	"smzdtz-server/pkg/util"

	"github.com/gin-gonic/gin"
)

const PasswordSaltLength = 6

// 创建管理员账号
func CreateAdminUser(name string, password string, email string) error {
	user := model.User{
		Name:     name,
		Password: password,
		Email:    email,
		IsAdmin:  1,
	}
	err := Create(user)
	if err != nil {
		return nil
	}
	return err
}

// 新增
func Create(user model.User) (err error) {
	user.Status = 1
	user.Salt = util.RandString(PasswordSaltLength)
	user.Password = util.Md5(user.Password + user.Salt)

	if err := model.Db.Create(&user).Error; err != nil {
		return err
	}
	return
}

// ValidateLogin 验证用户登录
func ValidateLogin(ctx *gin.Context) string {
	return "ValidateLogin"
	// username := ctx.QueryTrim("username")
	// password := ctx.QueryTrim("password")
	// json := utils.JsonResponse{}
	// if username == "" || password == "" {
	// 	return json.CommonFailure("用户名、密码不能为空")
	// }
	// userModel := new(models.User)
	// if !userModel.Match(username, password) {
	// 	return json.CommonFailure("用户名或密码错误")
	// }
	// loginLogModel := new(models.LoginLog)
	// loginLogModel.Username = userModel.Name
	// loginLogModel.Ip = ctx.RemoteAddr()
	// _, err := loginLogModel.Create()
	// if err != nil {
	// 	logger.Error("记录用户登录日志失败", err)
	// }

	// token, err := generateToken(userModel)
	// if err != nil {
	// 	logger.Errorf("生成jwt失败: %s", err)
	// 	return json.Failure(utils.AuthError, "认证失败")
	// }

	// return json.Success(utils.SuccessContent, map[string]interface{}{
	// 	"token":    token,
	// 	"uid":      userModel.Id,
	// 	"username": userModel.Name,
	// 	"is_admin": userModel.IsAdmin,
	// })
}
