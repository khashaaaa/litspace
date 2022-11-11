package psubcategory

import "github.com/gofiber/fiber/v2"

func PSubCategoryRouter(router fiber.Router) {

	psubcategory := router.Group("/psubcategory")

	psubcategory.Post("/store", Store)
	psubcategory.Get("/all", ShowAll)
	psubcategory.Get("/:mark", ShowSingle)
	psubcategory.Patch("/:mark", UpdateSingle)
	psubcategory.Delete("/:mark", DeleteSingle)
}
