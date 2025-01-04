package usecase

import (
	"bootstrap/domain"
	"bootstrap/usecase/task"
	"log"
)

const DELETE string = "delete_logs"
const UPDATE string = "update_database"

func (t taskUseCase) Task(taskInterface domain.TaskDomainInterface) string {

	taskStrategy := task.NewStrategy()
	taskStrategy.Register(DELETE, task.DeleteLogs{})
	taskStrategy.Register(UPDATE, task.UpdateDataBase{})

	run, err := taskStrategy.Get(taskInterface.GetAction())
	if err != nil {
		log.Panic("Strategy not found")
	}
	return run.Run(taskInterface.GetParameters())
}
