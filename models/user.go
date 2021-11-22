package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey;default:uuid_generate_v3()"`
	Name     string `json:"name" gorm:"not null;unique;size:32;comment:用户名"`
	Password string `json:"password" gorm:"not null;size:32;comment:密码"`
	Salt     string `json:"salt" gorm:"size:6;not null;comment:加盐"`
	Email    string `json:"email" gorm:"size:50;default:'';unique;comment:邮箱"`
	IsAdmin  int8   `json:"isAdmin" gorm:"size:1;default:0;comment:是否是管理员 1:管理员 0:普通用户"`
	Status   int8   `json:"status" gorm:"size:1;default:0;comment:1: 正常 0:禁用"`
	ModelTime
}
