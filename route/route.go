package route

import (
	"github.com/domjeff/golang-auth/controller"

	"github.com/gofiber/fiber"
)

func SetUp(app *fiber.App) {
	app.Get("/", controller.Hello)
	app.Post("api/register", controller.Register)
	app.Post("api/login", controller.Login)
	app.Listen(":8080")
}
