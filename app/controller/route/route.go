package route

import (
	controller "bootstrap/controller/task"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, controller controller.TaskControllerInterface) {
	r.POST("/task", controller.RunTask)
}
