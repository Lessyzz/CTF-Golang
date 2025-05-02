package main

import (
	"ctf/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	routes.SetupRoutes(app)
}

func main() {
	app := fiber.New()
	app.Static("/uploads", "./uploads")
	app.Static("/", "./public")

	SetupRoutes(app)

	app.Listen(":80")
}
