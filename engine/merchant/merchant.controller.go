package merchant

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khashaaaa/litspace/config"
)

func Register(context *fiber.Ctx) error {

	connector := config.InitConn()

	merchant := new(Merchant)

	parseErr := context.BodyParser(merchant)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	created, createErr := connector.Query("INSERT INTO merchant(founder, entity_name, email, mobile) VALUES($1, $2, $3, $4)", merchant.Founder, merchant.EntityName, merchant.Email, merchant.Mobile)

	if createErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": createErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "Худалдаачин бүртгэгдлээ",
		"data":    created,
	})
}

func ShowAll(context *fiber.Ctx) error {

	connector := config.InitConn()

	var merchantz = []Merchant{}
	var merchant Merchant

	rows, rowErr := connector.Query("SELECT mark, founder, entity_name, email, mobile, address, origin_country, buy_dest, sell_dest, in_status, type, created, updated FROM merchant")

	if rowErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": rowErr.Error(),
			"data":    nil,
		})
	}

	for rows.Next() {

		loopErr := rows.Scan(
			&merchant.Mark,
			&merchant.Founder,
			&merchant.EntityName,
			&merchant.Email,
			&merchant.Mobile,
			&merchant.Address,
			&merchant.OriginCountry,
			&merchant.BuyDest,
			&merchant.SellDest,
			&merchant.InStatus,
			&merchant.Type,
			&merchant.Created,
			&merchant.Updated,
		)

		if loopErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": loopErr.Error(),
				"data":    nil,
			})
		}

		merchantz = append(merchantz, merchant)
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Худалдаачдын жагсаалт",
		"data":    merchantz,
	})
}

func ShowSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	var merchant Merchant

	param := context.Params("id")

	query, queryErr := connector.Query("SELECT mark, founder, entity_name, email, mobile, address, origin_country, buy_dest, sell_dest, in_status, type, created, updated FROM merchant WHERE mark=$1", param)

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	for query.Next() {

		findErr := query.Scan(
			&merchant.Mark,
			&merchant.Founder,
			&merchant.EntityName,
			&merchant.Email,
			&merchant.Mobile,
			&merchant.Address,
			&merchant.OriginCountry,
			&merchant.BuyDest,
			&merchant.SellDest,
			&merchant.InStatus,
			&merchant.Type,
			&merchant.Created,
			&merchant.Updated,
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
		"message": "Хэрэглэгч",
		"data":    merchant,
	})
}

func UpdateSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("id")

	merchant := new(Merchant)

	parseErr := context.BodyParser(merchant)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE merchant SET founder=$1, entity_name=$2, email=$3, mobile=$4, address=$5, origin_country=$6, buy_dest=$7, sell_dest=$8, updated=$9 WHERE mark=$10")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&merchant.Founder,
		&merchant.EntityName,
		&merchant.Email,
		&merchant.Mobile,
		&merchant.Address,
		&merchant.OriginCountry,
		&merchant.BuyDest,
		&merchant.SellDest,
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
		"message": "Хэрэглэгчийн мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func UpdateType(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("id")

	merchant := new(Merchant)

	parseErr := context.BodyParser(merchant)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE merchant SET type=$1 WHERE mark=$2")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&merchant.Type,
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
		"message": "Худалдагчийн мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func DeleteSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("id")

	query, queryErr := connector.Prepare("DELETE FROM merchant WHERE mark=$1")

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
		"message": "Худалдагч устгагдлаа",
		"data":    exec,
	})
}