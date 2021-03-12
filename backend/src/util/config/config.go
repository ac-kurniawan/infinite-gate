package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
)

type PostgresConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type Config struct {
	PGConfig PostgresConnection
}

func (c *Config) ReadFromDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port,_ := strconv.Atoi(os.Getenv("PG_DB_PORT"))
	pgConnection := PostgresConnection{
		Name: os.Getenv("PG_DB_NAME"),
		Username: os.Getenv("PG_DB_USERNAME"),
		Password: os.Getenv("PG_DB_PASSWORD"),
		Host: os.Getenv("PG_DB_HOST"),
		Port: port,
	}

	c.PGConfig = pgConnection
}