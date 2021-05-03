package route

import (
	"github.com/domjeff/golang-auth/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Get("/", controller.Hello)
	app.Post("api/register", controller.Register)
	app.Post("api/login", controller.Login)
	app.Get("api/user", controller.User)
	app.Post("api/logout", controller.Logout)
	app.Listen(":8080")
}
