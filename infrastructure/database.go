package infrastructure

import (
	"fmt"
	"todo/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", configs.GetConfig().GetDbHost())
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	database = db
	return database
}

func GetDB() *gorm.DB {
	return database
}
