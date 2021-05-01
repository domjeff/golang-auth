package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	cred := os.Getenv("databaseuser") + ":" + os.Getenv("dabasepassword") + "@/golang-users"
	// fmt.Println(cred)
	_, err := gorm.Open(mysql.Open(cred), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
}
