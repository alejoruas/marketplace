package main

import (
	approuter "marketplace/infrastructure/approuter"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	approuter.StartRouter(router)
	router.Run(":8080")
}
