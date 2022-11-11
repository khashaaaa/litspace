package provider

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khashaaaa/litspace/config"
)

func Register(context *fiber.Ctx) error {

	connector := config.InitConn()

	provider := new(Provider)

	parseErr := context.BodyParser(provider)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	created, createErr := connector.Query("INSERT INTO provider(founder, entity_name, email, mobile) VALUES($1, $2, $3, $4)", provider.Founder, provider.EntityName, provider.Email, provider.Mobile)

	if createErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": createErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "Гүйцэтгэгч бүртгэгдлээ",
		"data":    created,
	})
}

func ShowAll(context *fiber.Ctx) error {

	connector := config.InitConn()

	var providerz []Provider
	var provider Provider

	rows, rowErr := connector.Query("SELECT mark, founder, entity_name, email, mobile, address, origin_country, in_status, type, created, updated FROM provider")

	if rowErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": rowErr.Error(),
			"data":    nil,
		})
	}

	for rows.Next() {

		loopErr := rows.Scan(
			&provider.Mark,
			&provider.Founder,
			&provider.EntityName,
			&provider.Email,
			&provider.Mobile,
			&provider.Address,
			&provider.OriginCountry,
			&provider.InStatus,
			&provider.Type,
			&provider.Created,
			&provider.Updated,
		)

		if loopErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": loopErr.Error(),
				"data":    nil,
			})
		}

		providerz = append(providerz, provider)
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Гүйцэтгэгчдийн жагсаалт",
		"data":    providerz,
	})
}

func ShowSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	var provider Provider

	param := context.Params("mark")

	query, queryErr := connector.Query("SELECT mark, founder, entity_name, email, mobile, address, origin_country, in_status, type, created, updated FROM provider WHERE mark=$1", param)

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	for query.Next() {

		findErr := query.Scan(
			&provider.Mark,
			&provider.Founder,
			&provider.EntityName,
			&provider.Email,
			&provider.Mobile,
			&provider.Address,
			&provider.OriginCountry,
			&provider.InStatus,
			&provider.Type,
			&provider.Created,
			&provider.Updated,
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
		"message": "Гүйцэтгэгч",
		"data":    provider,
	})
}

func UpdateSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	provider := new(Provider)

	parseErr := context.BodyParser(provider)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE provider SET founder=$1, entity_name=$2, email=$3, mobile=$4, address=$5, origin_country=$6, updated=$7 WHERE mark=$8")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&provider.Founder,
		&provider.EntityName,
		&provider.Email,
		&provider.Mobile,
		&provider.Address,
		&provider.OriginCountry,
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
		"message": "Гүйцэтгэгчийн мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func UpdateType(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	provider := new(Provider)

	parseErr := context.BodyParser(provider)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE provider SET type=$1 WHERE mark=$2")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&provider.Type,
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
		"message": "Гүйцэтгэгчийн мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func UpdateStatus(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	provider := new(Provider)

	parseErr := context.BodyParser(provider)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE provider SET in_status=$1 WHERE mark=$2")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&provider.InStatus,
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
		"message": "Гүйцэтгэгчийн мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func DeleteSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	query, queryErr := connector.Prepare("DELETE FROM provider WHERE mark=$1")

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
		"message": "Гүйцэтгэгч устгагдлаа",
		"data":    exec,
	})
}
