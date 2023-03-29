package main

import (
	"github.com/gofiber/fiber/v2"
	"webserser/routes"
)

func main() {
	// create a fiber application
	var app *fiber.App = fiber.New()
	// add a request handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("{}")
	})
	routes.SetupRoutes(app)

	// start the application at port 3000
	app.Listen(":3000")
}
