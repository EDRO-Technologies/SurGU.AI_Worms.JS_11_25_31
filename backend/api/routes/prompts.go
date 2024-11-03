package routes

import (
	"github.com/gofiber/fiber/v2"
	"worm-pack/api/controllers"
	mw "worm-pack/api/middleware"
	C "worm-pack/constants"
)

func SetupProductsRoutes(router fiber.Router) {
	router.Post("/prompt", mw.RateLimit(C.Tier7, 0), controllers.PostPrompt)

	router.Post("/legendaryPrompt", mw.RateLimit(C.Tier7, 0), controllers.PingPong)

}
