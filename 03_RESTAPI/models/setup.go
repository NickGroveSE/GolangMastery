package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := "root"
	PASS := "password"
	HOST := "127.0.0.1"
	PORT := "3306"
	DBNAME := "untappdCRUD"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
