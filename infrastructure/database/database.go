package database

import (
	"fmt"
	"todo/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", config.GetConfig().GetDbHost())
	if err != nil {
		fmt.Println("database init error: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	database = db
	return database
}
