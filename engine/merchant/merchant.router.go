package merchant

import "github.com/gofiber/fiber/v2"

func MerchantRouter(router fiber.Router) {

	merchant := router.Group("/merchant")

	merchant.Post("/register", Register)
	merchant.Get("/all", ShowAll)
	merchant.Get("/:id", ShowSingle)
	merchant.Patch("/:id", UpdateSingle)
	merchant.Patch("/:id/type", UpdateType)
	merchant.Delete("/:id", DeleteSingle)
}
