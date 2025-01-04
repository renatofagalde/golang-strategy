package domain

import "time"

type taskDomain struct {
	createdAt  time.Time
	action     string
	parameters string
	typeUsed   string
}

func (t *taskDomain) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *taskDomain) GetAction() string {
	return t.action
}

func (t *taskDomain) GetParameters() string {
	return t.parameters
}

func (t *taskDomain) GetType() string {
	return t.typeUsed
}
