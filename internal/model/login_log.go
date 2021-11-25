package model

// 用户登录日志
type LoginLog struct {
	ModelId
	Username string `json:"username" gorm:"size(32);not null"`
	Ip       string `json:"ip" gorm:"size(15);not null"`
	ModelTime
}
