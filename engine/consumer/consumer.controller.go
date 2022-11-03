package consumer

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khashaaaa/litspace/config"
	"golang.org/x/crypto/bcrypt"
)

func Hash(pazzword string) (string, error) {
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(pazzword), bcrypt.DefaultCost)
	return string(hash), hashErr
}

func Compare(pazz string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pazz), []byte(hashed))
	return err
}

func Register(context *fiber.Ctx) error {

	connector := config.InitConn()

	consumer := new(Consumer)

	parseErr := context.BodyParser(consumer)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	hashed, hashErr := Hash(consumer.Pass)

	if hashErr != nil {
		return context.Status(http.StatusNotAcceptable).JSON(&fiber.Map{
			"status":  http.StatusNotAcceptable,
			"message": hashErr.Error(),
			"data":    nil,
		})
	}

	created, createErr := connector.Query("INSERT INTO consumer(first_name, last_name, email, pass) VALUES($1, $2, $3, $4)", consumer.FirstName, consumer.LastName, consumer.Email, hashed)

	if createErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": createErr.Error(),
			"data":    nil,
		})
	}

	return context.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "Хэрэглэгч бүртгэгдлээ",
		"data":    created,
	})
}

func Login(context *fiber.Ctx) error {

	connector := config.InitConn()

	var auth = Auth{}

	parseErr1 := context.BodyParser(&auth)

	if parseErr1 != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr1.Error(),
			"data":    nil,
		})
	}

	found, foundErr := connector.Query("SELECT mark, first_name, last_name, email, mobile, origin_country, pass, type, created, updated FROM consumer WHERE email=$1", auth.Email)

	if foundErr != nil {
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"status":  http.StatusNotFound,
			"message": foundErr.Error(),
			"data":    nil,
		})
	}

	var consumer = Consumer{}

	parseErr2 := context.BodyParser(&consumer)

	if parseErr2 != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr2.Error(),
			"data":    nil,
		})
	}

	for found.Next() {

		findErr := found.Scan(
			&consumer.Mark,
			&consumer.FirstName,
			&consumer.LastName,
			&consumer.Email,
			&consumer.Mobile,
			&consumer.OriginCountry,
			&consumer.Pass,
			&consumer.Type,
			&consumer.Created,
			&consumer.Updated,
		)

		if findErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": findErr.Error(),
				"data":    nil,
			})
		}
	}

	compareError := Compare(consumer.Pass, auth.Pass)

	if compareError != nil {
		return context.Status(http.StatusForbidden).JSON(&fiber.Map{
			"status":  http.StatusForbidden,
			"message": compareError.Error(),
			"data":    nil,
		})
	}

	consumer.Pass = auth.Pass

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Хэрэглэгч",
		"data":    consumer,
	})
}

func ShowAll(context *fiber.Ctx) error {

	connector := config.InitConn()

	var consumerz = []Consumer{}
	var consumer Consumer

	rows, rowErr := connector.Query("SELECT mark, first_name, last_name, email, mobile, origin_country, pass, type, created, updated FROM consumer")

	if rowErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": rowErr.Error(),
			"data":    nil,
		})
	}

	for rows.Next() {

		loopErr := rows.Scan(
			&consumer.Mark,
			&consumer.FirstName,
			&consumer.LastName,
			&consumer.Email,
			&consumer.Mobile,
			&consumer.OriginCountry,
			&consumer.Pass,
			&consumer.Type,
			&consumer.Created,
			&consumer.Updated,
		)

		if loopErr != nil {
			return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": loopErr.Error(),
				"data":    nil,
			})
		}

		consumerz = append(consumerz, consumer)
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "Хэрэглэгчийн жагсаалт",
		"data":    consumerz,
	})
}

func ShowSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	var consumer = Consumer{}

	param := context.Params("id")

	query, queryErr := connector.Query("SELECT mark, first_name, last_name, email, mobile, origin_country, pass, type, created, updated FROM consumer WHERE mark=$1", param)

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	for query.Next() {

		findErr := query.Scan(
			&consumer.Mark,
			&consumer.FirstName,
			&consumer.LastName,
			&consumer.Email,
			&consumer.Mobile,
			&consumer.OriginCountry,
			&consumer.Pass,
			&consumer.Type,
			&consumer.Created,
			&consumer.Updated,
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
		"data":    consumer,
	})
}

func UpdateSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("id")

	consumer := new(Consumer)

	parseErr := context.BodyParser(consumer)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE consumer SET first_name=$1, last_name=$2, email=$3, mobile=$4, origin_country=$5, updated=$6 WHERE mark=$7")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&consumer.FirstName,
		&consumer.LastName,
		&consumer.Email,
		&consumer.Mobile,
		&consumer.OriginCountry,
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

func UpdatePass(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("id")

	consumer := new(Consumer)

	parseErr := context.BodyParser(consumer)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE consumer SET pass=$1 WHERE mark=$2")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&consumer.Pass,
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

	consumer := new(Consumer)

	parseErr := context.BodyParser(consumer)

	if parseErr != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  http.StatusBadRequest,
			"message": parseErr.Error(),
			"data":    nil,
		})
	}

	query, queryErr := connector.Prepare("UPDATE consumer SET type=$1 WHERE mark=$2")

	if queryErr != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": queryErr.Error(),
			"data":    nil,
		})
	}

	exec, execErr := query.Exec(
		&consumer.Type,
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

func DeleteSingle(context *fiber.Ctx) error {

	connector := config.InitConn()

	param := context.Params("id")

	query, queryErr := connector.Prepare("DELETE FROM consumer WHERE mark=$1")

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
		"message": "Хэрэглэгч устгагдлаа",
		"data":    exec,
	})
}
