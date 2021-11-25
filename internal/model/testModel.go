package model

import "fmt"

type Test struct {
	Id      int
	Testcol string `gorm:"column:testcol"`
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// }

func (this *Test) Insert() (id int, err error) {
	fmt.Println("11")
	// result := Mysql.DB.Create(&this)
	// id = this.Id
	// if result.Error != nil {
	// 	err = result.Error
	// 	return
	// }
	return 1, nil
}
