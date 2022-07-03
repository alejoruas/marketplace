package main

import (
	"context"
	"log"
	approuter "marketplace/infrastructure/approuter"
	"marketplace/infrastructure/database"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	sql, err := database.NewDBSQL()

	if err != nil {
		log.Println(err)
		return
	}

	router := gin.Default()
	approuter.StartRouter(router, sql)

	ginLambda = ginadapter.New(router)
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
