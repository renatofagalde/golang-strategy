package task

import "fmt"

type DeleteLogs struct {
}

func (d DeleteLogs) Run(logs string) string {
	fmt.Print(logs)
	return "deleted"
}
