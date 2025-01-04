package main

import (
	"bootstrap/controller/route"
	controller "bootstrap/controller/task"
	"bootstrap/usecase"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	service := usecase.NewTaskUsecase()
	taskController := controller.NewControllerInterface(service)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	route.InitRoutes(&router.RouterGroup, taskController)
	ginLambda = ginadapter.New(router)

}

func main() {
	lambda.Start(ginLambda.ProxyWithContext)
}
