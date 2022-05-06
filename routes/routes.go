package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// Welcome
	app.Get("/", controllers.Welcome)

	// User
	app.Post("/api/users", controllers.RegisterUser)
}
