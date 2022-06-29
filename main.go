package main

import (
	"fmt"
	approuter "marketplace/infrastructure/approuter"
	db "marketplace/infrastructure/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	sql, err := db.NewDBSQL()

	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	approuter.StartRouter(router, sql)
	router.Run(os.Getenv("APP_PORT"))
}
