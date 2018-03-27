package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "root:neogov@123@/todo")
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
