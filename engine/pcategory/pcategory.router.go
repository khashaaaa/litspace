package pcategory

import "github.com/gofiber/fiber/v2"

func PCategoryRouter(router fiber.Router) {

	pcategory := router.Group("/pcategory")

	pcategory.Post("/store", Store)
	pcategory.Get("/all", ShowAll)
	pcategory.Get("/:mark", ShowSingle)
	pcategory.Patch("/:mark", UpdateSingle)
	pcategory.Delete("/:mark", DeleteSingle)
}
