package main

import (
	"github.com/domjeff/golang-auth/database"
	"github.com/domjeff/golang-auth/route"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()
	app := fiber.New()
	route.SetUp(app)
	app.Listen(":8080")

}
