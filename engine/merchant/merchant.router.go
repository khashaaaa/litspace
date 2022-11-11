package merchant

import "github.com/gofiber/fiber/v2"

func MerchantRouter(router fiber.Router) {

	merchant := router.Group("/merchant")

	merchant.Post("/register", Register)
	merchant.Get("/all", ShowAll)
	merchant.Get("/:mark", ShowSingle)
	merchant.Patch("/:mark", UpdateSingle)
	merchant.Patch("/:mark/type", UpdateType)
	merchant.Patch("/:mark/status", UpdateStatus)
	merchant.Delete("/:mark", DeleteSingle)
}
