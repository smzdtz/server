package models

import "errors"

type Migration struct{}

// 首次安装, 创建数据库表
func (migration *Migration) Install() error {
	user := &User{}
	tables := []interface{}{user}
	for _, table := range tables {
		exist := Db.Migrator().HasTable(table)
		if exist {
			return errors.New("数据表已存在")
		}
		err := Db.Migrator().CreateTable(table)
		if err != nil {
			return err
		}
	}
	// setting.InitBasicField()
	return nil
}
