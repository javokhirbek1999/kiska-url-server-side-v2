package main

import (
	"core/database"
	"core/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello Kiska URL")
}

func main() {

	app := fiber.New()

	database.ConnectDB()

	routes.Routes(app)

	log.Fatal(app.Listen(":8000"))

}
