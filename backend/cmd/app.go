package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"worm-pack/api/routes"
	_ "worm-pack/docs"
	H "worm-pack/handler"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	worms.js

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		0.0.0.0:70
//	@BasePath	/

func InitApp() *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: H.ErrorHandler,
		},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, HEAD",
	}))

	app.Get("/swagger/*", swagger.New(swagger.ConfigDefault))

	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	app.Use(requestid.New())

	app.Use(logger.New())

	routes.SetupRoutes(app)

	return app
}
