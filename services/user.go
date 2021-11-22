package services

import (
	"smzdtz-server/models"
	"smzdtz-server/utils"
)

const PasswordSaltLength = 6

// 创建管理员账号
func createAdminUser(form models.InstallForm) error {
	user := new(models.User)
	user.Name = form.AdminUsername
	user.Password = form.AdminPassword
	user.Email = form.AdminEmail
	user.IsAdmin = 1
	_, err := Create()

	return err
}

// 新增
func Create() (insertId int, err error) {
	user := models.User{}
	user.Status = 1
	user.Salt = utils.RandString(PasswordSaltLength)
	user.Password = utils.Md5(user.Password + user.Salt)

	models.Db.Create(&user)

	return
}
