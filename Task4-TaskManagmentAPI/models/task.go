package models

type Task struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Start    string `json:"start"`
	Duration string `json:"duration"`
}
