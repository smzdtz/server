package models

import "gorm.io/gorm"

type Status int8

var TablePrefix = ""
var Db *gorm.DB

const (
	Disabled Status = 0 // 禁用
	Failure  Status = 0 // 失败
	Enabled  Status = 1 // 启用
	Running  Status = 1 // 运行中
	Finish   Status = 2 // 完成
	Cancel   Status = 3 // 取消
)
