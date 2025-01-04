package task

import "fmt"

type UpdateDataBase struct {
}

func (d UpdateDataBase) Run(stmt string) string {

	msg := fmt.Sprintf("Update database with: %s ", stmt)
	fmt.Print(msg)
	return "Data base updated"
}
