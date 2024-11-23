package config

// this package is connect to data base and save data in this db
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// func to connection to the MySQL database
	d, err := gorm.Open("mysql", "belok:admin/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	// get access to the current database connection
	return db
}
