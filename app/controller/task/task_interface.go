package controller

import (
	"bootstrap/usecase"
	"github.com/gin-gonic/gin"
)

type TaskControllerInterface interface {
	RunTask(c *gin.Context)
}

func NewControllerInterface(service usecase.TaskUsecase) TaskControllerInterface {
	return &taskController{service}
}

type taskController struct {
	service usecase.TaskUsecase
}
