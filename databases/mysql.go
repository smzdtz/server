package Mysql

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// var err error
	// DB, err = gorm.Open("mysql", "wuyu:MIDSUMMERfish0@/gin?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Printf("mysql connect error %v", err)
	// }
	// if DB.Error != nil {
	// 	fmt.Printf("database error %v", DB.Error)
	// }
	fmt.Printf("database error %v", DB.Error)
}
