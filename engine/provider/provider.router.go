package provider

import "github.com/gofiber/fiber/v2"

func ProviderRouter(router fiber.Router) {

	provider := router.Group("/provider")

	provider.Post("/register", Register)
	provider.Get("/all", ShowAll)
	provider.Get("/:mark", ShowSingle)
	provider.Patch("/:mark", UpdateSingle)
	provider.Patch("/:mark/type", UpdateType)
	provider.Patch("/:mark/status", UpdateStatus)
	provider.Delete("/:mark", DeleteSingle)
}
