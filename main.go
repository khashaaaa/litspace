package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/khashaaaa/litspace/config"
	"github.com/khashaaaa/litspace/engine/consumer"
	"github.com/khashaaaa/litspace/engine/merchant"
)

func main() {
	prog := fiber.New()
	prog.Use(cors.New())
	config.InitConn()

	consumer.ConsumerRouter(prog)
	merchant.MerchantRouter(prog)

	prog.Listen(":8383")
}
