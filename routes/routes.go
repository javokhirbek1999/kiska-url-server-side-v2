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
	app.Post("api/auth/login", controllers.Login)
	app.Get("api/auth/user", controllers.User)
	app.Post("api/auth/logout", controllers.Logout)

	// Urls
	app.Post("api/urls/shorten", controllers.ShortenURL)
	app.Get("/:urlHash", controllers.Redirect)
}
