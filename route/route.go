package route

import (
	"github.com/domjeff/golang-auth/controller"
	"github.com/gofiber/fiber"
)

func SetUp(app *fiber.App) {
	app.Post("api/register", controller.Register)
	app.Get("/", controller.Hello)

	app.Listen(":8080")
}
