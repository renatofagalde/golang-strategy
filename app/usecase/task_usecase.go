package usecase

import "bootstrap/domain"

func NewTaskUsecase() TaskUsecase {
	return &taskUseCase{}
}

type taskUseCase struct {
}

type TaskUsecase interface {
	Task(taskInterface domain.TaskDomainInterface) string
}
