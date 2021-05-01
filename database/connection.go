package database

import (
	"log"
	"os"

	"github.com/domjeff/golang-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cred := os.Getenv("databaseuser") + ":" + os.Getenv("dabasepassword") + "@/golang-users"

	connection, err := gorm.Open(mysql.Open(cred), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
