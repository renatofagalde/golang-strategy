package controller

import (
	"bootstrap/controller/task/model"
	"bootstrap/domain"
	"bootstrap/view"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *taskController) RunTask(c *gin.Context) {

	var taskRequest model.TaskRequest

	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, "bad_request")
		return
	}

	result := t.service.Task(domain.NewTaskDomain(taskRequest.Action, taskRequest.Parameter))
	fmt.Println("result -> ", result)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))

}
