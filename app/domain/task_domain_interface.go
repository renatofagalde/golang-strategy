package domain

import "time"

// Constructor in golang
func NewTaskDomain(action string, parameters string) TaskDomainInterface {
	return &taskDomain{
		createdAt:  time.Now(),
		action:     action,
		parameters: parameters,
	}
}

type TaskDomainInterface interface {
	GetCreatedAt() time.Time
	GetAction() string
	GetParameters() string
	GetType() string
}
