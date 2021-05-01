package env

import (
	"log"

	"github.com/joho/godotenv"
)

func InitiateEnvVar() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("%v", err.Error())
	}
}
