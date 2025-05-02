package routes

import (
	"ctf/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/")
	route.Get("/", handlers.Index_GET)

	route.Get("/adminlogin", handlers.Login_GET)
	route.Post("/adminlogin", handlers.Login_POST)

	route.Get("/admin/dashboard", handlers.Dashboard_GET)
	route.Get("/admin/upload", handlers.Upload_GET)
	route.Post("/admin/upload", handlers.Upload_POST)

	route.Get("/privateimage", handlers.Image_GET)
	route.Get("/flag", handlers.Flag_GET)
	route.Post("/flag", handlers.Flag_POST)
}
