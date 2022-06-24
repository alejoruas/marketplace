package main

import (
	"fmt"
	approuter "marketplace/infrastructure/approuter"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	approuter.StartRouter(router)

	fmt.Println(os.Getenv("APP_PORT"))

	router.Run(os.Getenv("APP_PORT"))
}
