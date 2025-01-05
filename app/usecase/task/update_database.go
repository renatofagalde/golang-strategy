package task

import "fmt"

type UpdateDataBase struct {
}

func (d UpdateDataBase) Run(stmt string) string {

	msg := fmt.Sprintf("Database has been updated with: %s", stmt)
	return msg
}
