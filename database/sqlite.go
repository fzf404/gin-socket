/*
 * @Author: fzf404
 * @Date: 2021-09-21 22:29:46
 * @LastEditTime: 2021-09-22 19:48:10
 * @Description: 初始化Sqlite
 */
package database

import (
	"gin-socket/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

/**
 * @description: 初始化Sqlite
 */
func InitSQLite() {
	db, err := gorm.Open(sqlite.Open("./database/database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.WebApp{})
	DB = db
}
