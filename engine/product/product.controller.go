package product

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khashaaaa/litspace/config"
	"github.com/lib/pq"
)

func Store(context *fiber.Ctx) error {

	connector := config.InitConn()

	product := new(Product)

	parseErr := context.BodyParser(product)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	newAttr, _ := json.Marshal(product.Attrs)

	created, createErr := connector.Query("INSERT INTO product(merchant, category, name, descr, price, stock, attrs, image_paths) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", product.Merchant, product.Category, product.Name, product.Descr, product.Price, product.Stock, newAttr, pq.Array(product.ImagePaths))

	if createErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": createErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "Бараа хадгалагдлаа",
		"data":    created,
	})
}

func ShowAll(context *fiber.Ctx) error {

	connector := config.InitConn()

	var productz []Product
	var product Product

	rows, rowErr := connector.Query("SELECT mark, merchant, category, name, descr, price, stock, attrs, image_paths, created, updated FROM product")

	if rowErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": rowErr.Error(),
			"data":    nil,
		})
	}

	for rows.Next() {

		loopErr := rows.Scan(
			&product.Mark,
			&product.Merchant,
			&product.Category,
			&product.Name,
			&product.Descr,
			&product.Price,
			&product.Stock,
			&product.Attrs,
			&product.ImagePaths,
			&product.Created,
			&product.Updated,
		)

		if loopErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": loopErr.Error(),
				"data":    nil,
			})
		}

		productz = append(productz, product)
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Барааны жагсаалт",
		"data":    productz,
	})
}

func ShowSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	var product Product

	param := context.Params("mark")

	query, queryErr := connector.Query("SELECT mark, merchant, category, name, descr, price, stock, attrs, image_paths, created, updated FROM product WHERE mark=$1", param)

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	for query.Next() {

		findErr := query.Scan(
			&product.Mark,
			&product.Merchant,
			&product.Category,
			&product.Name,
			&product.Descr,
			&product.Price,
			&product.Stock,
			&product.Attrs,
			&product.ImagePaths,
			&product.Created,
			&product.Updated,
		)

		if findErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": findErr.Error(),
				"data":    nil,
			})
		}
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Бараа",
		"data":    product,
	})
}

func UpdateSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	product := new(Product)

	parseErr := context.BodyParser(product)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE product SET category=$1, name=$2, descr=$3, price=$4, stock=$5, attrs=$6, image_paths=$7, updated=$8 WHERE mark=$9")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&product.Category,
		&product.Name,
		&product.Descr,
		&product.Price,
		&product.Stock,
		&product.Attrs,
		&product.ImagePaths,
		time.Now(),
		param,
	)

	if execErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": execErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Барааны мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func DeleteSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	query, queryErr := connector.Prepare("DELETE FROM product WHERE mark=$1")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(param)

	if execErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": execErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Бараа устгагдлаа",
		"data":    exec,
	})
}
