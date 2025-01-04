package view

import (
	"bootstrap/controller/task/model"
)

func ConvertDomainToResponse(result string) model.TaskResponse {
	return model.TaskResponse{Result: result}
}
