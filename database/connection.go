package database

import (
	"github.com/kibo/e-wallet/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {
	connection, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/golang_test?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	} 

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.History{})
}