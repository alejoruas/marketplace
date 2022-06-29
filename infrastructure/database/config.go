package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type configdb struct {
	Host     string
	Driver   string
	Port     string
	Database string
	User     string
	Password string
}

func CreatePostgresConfig() *configdb {
	var config = configdb{
		Host:     os.Getenv("DB_HOST"),
		Driver:   os.Getenv("DB_DRIVER"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	config.showConfig()

	return &config
}

func (c configdb) showConfig() {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
