package model

type TaskRequest struct {
	Action    string `json:"action" binding:"required,min=3"`
	Parameter string `json:"parameter" binding:"required,min=3"`
}
