package main

import (
	"context"
	"log"
	approuter "marketplace/infrastructure/approuter"
	"marketplace/infrastructure/database"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	sql, err := database.NewDBSQL()

	if err != nil {
		log.Println(err)
		return
	}

	router := gin.Default()
	approuter.StartRouter(router, sql)
	router.Run(os.Getenv("APP_PORT"))

	//ginLambda = ginadapter.New(router)
	//lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
