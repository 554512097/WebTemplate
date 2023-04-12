package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var GloableDB *gorm.DB

func InitDB() error {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/goon?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	GloableDB = db
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// 启用Logger，显示详细日志
	db.LogMode(true)

	//创建表
	db.CreateTable(&User{})
	return nil
}
