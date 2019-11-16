package api_manager

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func OpenDb() {
	temp, err := gorm.Open("sqlite3", "a_api_manager.db")
	if err != nil {
		panic(err)
	}
	temp.AutoMigrate(Api{})
	temp.AutoMigrate(User{})
	temp.AutoMigrate(ApiCollection{})
	db = temp
	db.LogMode(true)
}
