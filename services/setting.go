package services

import (
	"smzdtz-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 创建临时数据库连接
func createTmpDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.temp.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

func TestDbConnection(form models.InstallForm) error {
	db, _ := createTmpDb()
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	return err
}

// 创建Db
func CreateDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
