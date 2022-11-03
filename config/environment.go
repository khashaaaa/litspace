package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnv(key string) string {

	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Println("Cannot find env file")
	}

	return os.Getenv(key)
}
