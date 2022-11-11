package ssubcategory

import "github.com/gofiber/fiber/v2"

func SSubCategoryRouter(router fiber.Router) {

	ssubcategory := router.Group("/ssubcategory")

	ssubcategory.Post("/store", Store)
	ssubcategory.Get("/all", ShowAll)
	ssubcategory.Get("/:mark", ShowSingle)
	ssubcategory.Patch("/:mark", UpdateSingle)
	ssubcategory.Delete("/:mark", DeleteSingle)
}
