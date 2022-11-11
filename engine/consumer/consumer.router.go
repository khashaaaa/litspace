package consumer

import "github.com/gofiber/fiber/v2"

func ConsumerRouter(router fiber.Router) {

	consumer := router.Group("/consumer")

	consumer.Post("/register", Register)
	consumer.Post("/login", Login)
	consumer.Get("/all", ShowAll)
	consumer.Get("/:mark", ShowSingle)
	consumer.Patch("/:mark", UpdateSingle)
	consumer.Patch("/:mark/pass", UpdatePass)
	consumer.Patch("/:mark/type", UpdateType)
	consumer.Delete("/:mark", DeleteSingle)
}
