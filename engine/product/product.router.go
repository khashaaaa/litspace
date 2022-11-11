package product

import "github.com/gofiber/fiber/v2"

func ProductRouter(router fiber.Router) {

	product := router.Group("/product")

	product.Post("/store", Store)
	product.Get("/all", ShowAll)
	product.Get("/:mark", ShowSingle)
	product.Patch("/:mark", UpdateSingle)
	product.Delete("/:mark", DeleteSingle)
}
