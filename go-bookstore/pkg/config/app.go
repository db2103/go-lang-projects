package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			"root",
			"admin",
			"localhost",
			3306,
			"simplerest",
		),
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database!, %v", err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
