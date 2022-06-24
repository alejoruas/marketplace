package database

import (
	"os"
)

type configdb struct {
	host     string
	driver   string
	port     string
	database string
	user     string
	password string
}

func createPostgresConfig() *configdb {
	return &configdb{
		host:     os.Getenv("DB_HOST"),
		driver:   os.Getenv("DB_DRIVER"),
		port:     os.Getenv("DB_PORT"),
		database: os.Getenv("DB_NAME=project"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
	}
}
