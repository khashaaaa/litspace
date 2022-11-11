package scategory

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khashaaaa/litspace/config"
)

func Store(context *fiber.Ctx) error {

	connector := config.InitConn()

	data := new(SCategory)

	parseErr := context.BodyParser(data)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	created, createErr := connector.Query("INSERT INTO scategory(name) VALUES($1)", data.Name)

	if createErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": createErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "Категори бүртгэгдлээ",
		"data":    created,
	})
}

func ShowAll(context *fiber.Ctx) error {

	connector := config.InitConn()

	var dataz []SCategory
	var data SCategory

	rows, rowErr := connector.Query("SELECT mark, name, created, updated FROM scategory")

	if rowErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": rowErr.Error(),
			"data":    nil,
		})
	}

	for rows.Next() {

		loopErr := rows.Scan(
			&data.Mark,
			&data.Name,
			&data.Created,
			&data.Updated,
		)

		if loopErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": loopErr.Error(),
				"data":    nil,
			})
		}

		dataz = append(dataz, data)
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Категорийн жагсаалт",
		"data":    dataz,
	})
}

func ShowSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	var data SCategory

	param := context.Params("mark")

	query, queryErr := connector.Query("SELECT mark, name, created, updated FROM scategory WHERE mark=$1", param)

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	for query.Next() {

		findErr := query.Scan(
			&data.Mark,
			&data.Name,
			&data.Created,
			&data.Updated,
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
		"message": "Категори",
		"data":    data,
	})
}

func UpdateSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	data := new(SCategory)

	parseErr := context.BodyParser(data)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE scategory SET name=$1, updated=$2 WHERE mark=$3")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&data.Name,
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
		"message": "Категорийн мэдээлэл өөрчлөгдлөө",
		"data":    exec,
	})
}

func DeleteSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("mark")

	query, queryErr := connector.Prepare("DELETE FROM scategory WHERE mark=$1")

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
		"message": "Категори устгагдлаа",
		"data":    exec,
	})
}
