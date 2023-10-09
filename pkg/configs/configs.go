package configs

import (
	"fmt"
	"iudx_domain_specific_apis/pkg/logger"
	"os"
)

type Config struct {
	environment string
	dbUser      string
	dbPass      string
	dbHost      string
	dbPort      string
	dbName      string
	apiPort     string
	apiKey      string
}

var config Config

func Initialize() {
	config.environment = os.Getenv("GO_ENV")
	config.dbUser = os.Getenv("POSTGRES_USER")
	config.dbPass = os.Getenv("POSTGRES_PASSWORD")
	config.dbHost = os.Getenv("POSTGRES_HOST")
	config.dbPort = os.Getenv("POSTGRES_PORT")
	config.dbName = os.Getenv("POSTGRES_DB")
	config.apiPort = os.Getenv("API_PORT")
	config.apiKey = os.Getenv("API_AUTH_KEY")

	logger.Info.Println("Successfully loaded all config")
}

func GetDBConnStr() string {
	return config.getDBConnStr(config.dbHost, config.dbName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPass,
		dbhost,
		c.dbPort,
		dbname,
	)
}

func GetAPIPort() string {
	return ":" + config.apiPort
}

func GetEnvironment() string {
	return config.environment
}

func GetAPIKey() string {
	return config.apiKey
}
