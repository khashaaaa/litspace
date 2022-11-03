package consumer

import "github.com/gofiber/fiber/v2"

func ConsumerRouter(router fiber.Router) {

	consumer := router.Group("/consumer")

	consumer.Post("/register", Register)
	consumer.Post("/login", Login)
	consumer.Get("/all", ShowAll)
	consumer.Get("/:id", ShowSingle)
	consumer.Patch("/:id", UpdateSingle)
	consumer.Patch("/:id/pass", UpdatePass)
	consumer.Patch("/:id/type", UpdateType)
	consumer.Delete("/:id", DeleteSingle)
}
