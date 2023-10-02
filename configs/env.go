package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_DB_NAME  string
}

func GetEnv() Env {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	var envObj Env

	envObj.POSTGRES_DB_NAME = os.Getenv("POSTGRES_USER")
	envObj.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	envObj.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	envObj.POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")

	return envObj
}
