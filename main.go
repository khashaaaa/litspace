package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/khashaaaa/litspace/config"
	"github.com/khashaaaa/litspace/engine/consumer"
	"github.com/khashaaaa/litspace/engine/merchant"
	"github.com/khashaaaa/litspace/engine/pcategory"
	"github.com/khashaaaa/litspace/engine/product"
	"github.com/khashaaaa/litspace/engine/provider"
	"github.com/khashaaaa/litspace/engine/psubcategory"
	"github.com/khashaaaa/litspace/engine/scategory"
	"github.com/khashaaaa/litspace/engine/ssubcategory"
)

func main() {
	prog := fiber.New()
	prog.Use(cors.New())
	config.InitConn()

	consumer.ConsumerRouter(prog)
	merchant.MerchantRouter(prog)
	provider.ProviderRouter(prog)
	pcategory.PCategoryRouter(prog)
	psubcategory.PSubCategoryRouter(prog)
	scategory.SCategoryRouter(prog)
	ssubcategory.SSubCategoryRouter(prog)
	product.ProductRouter(prog)

	prog.Listen(":8383")
}
