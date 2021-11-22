package services

import (
	"smzdtz-server/models"
	"smzdtz-server/utils"
)

const PasswordSaltLength = 6

// 创建管理员账号
func CreateAdminUser(params models.InstallForm) error {
	user := models.User{
		Name:     params.AdminUsername,
		Password: params.AdminPassword,
		Email:    params.AdminEmail,
		IsAdmin:  1,
	}
	err := Create(user)
	if err != nil {
		return nil
	}
	return err
}

// 新增
func Create(user models.User) (err error) {
	user.Status = 1
	user.Salt = utils.RandString(PasswordSaltLength)
	user.Password = utils.Md5(user.Password + user.Salt)

	if err := models.Db.Create(&user).Error; err != nil {
		return err
	}
	return
}
