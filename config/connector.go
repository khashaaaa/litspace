package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = ReadEnv("PG_HOST")
	user     = ReadEnv("PG_USER")
	password = ReadEnv("PG_PWD")
	dbname   = ReadEnv("PG_DB")
	port     = ReadEnv("PG_PORT")
)

func InitConn() *sql.DB {

	ztring := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ulaanbaatar", host, user, password, dbname, port)
	connector, fault := sql.Open("postgres", ztring)

	if fault != nil {
		log.Println("Error occured to connection: ", fault.Error())
		return nil
	}

	log.Println("Database connection established")

	return connector
}
