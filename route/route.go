package route

import (
	"github.com/domjeff/golang-auth/controller"

	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Get("/", controller.Hello)
	app.Post("api/register", controller.Register)
	app.Post("api/login", controller.Login)
	app.Post("api/test", controller.Test)
	app.Listen(":8080")
}
