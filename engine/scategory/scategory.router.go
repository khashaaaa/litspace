package scategory

import "github.com/gofiber/fiber/v2"

func SCategoryRouter(router fiber.Router) {

	scategory := router.Group("/scategory")

	scategory.Post("/store", Store)
	scategory.Get("/all", ShowAll)
	scategory.Get("/:mark", ShowSingle)
	scategory.Patch("/:mark", UpdateSingle)
	scategory.Delete("/:mark", DeleteSingle)
}
