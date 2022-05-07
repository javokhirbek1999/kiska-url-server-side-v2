package routes

import (
	"core/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	// Welcome
	app.Get("/", controllers.Welcome)

	// User
	app.Post("/api/auth/users", controllers.RegisterUser)
	app.Post("api/auth/confirm/:token", controllers.EmailConfirmation)
}
