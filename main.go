package main

import (
	"github.com/domjeff/golang-auth/database"
	"github.com/domjeff/golang-auth/env"
	"github.com/domjeff/golang-auth/route"

	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	env.InitiateEnvVar()
	database.Connect()
	app := fiber.New()
	app.Use(
		cors.New(cors.Config{
			AllowCredentials: true,
		}),
	)
	route.SetUp(app)

	app.Listen(":8080")

}
