package routes

import (
	"github.com/gofiber/fiber/v2"
	"worm-pack/api/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Health)

	v1API := app.Group("/api/")

	SetupProductsRoutes(v1API)
}
