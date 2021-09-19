package config

import (
	"golang-crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	connectionString := "root:@tcp(localhost:3306)/golang_crud?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
