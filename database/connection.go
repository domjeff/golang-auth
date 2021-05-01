package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:inipassword@/golang-users"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
}
