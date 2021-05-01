package main

import (
	"github.com/domjeff/golang-auth/database"
	"github.com/domjeff/golang-auth/env"
	"github.com/domjeff/golang-auth/route"
	"github.com/gofiber/fiber"
)

func main() {
	env.InitiateEnvVar()
	database.Connect()
	app := fiber.New()
	route.SetUp(app)
	app.Listen(":8080")

}
