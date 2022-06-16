package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	localhost := "localhost"
	password := "dev123"
	username := "root"
	database := "crud_go"
	port := 3306

	url := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, localhost, port, database)

	con, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return con
}

func AutoMigrate() {
	DB().AutoMigrate(&Entity{})
}
