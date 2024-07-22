package models

type Task struct {
	ID          uint   `json:"id" example:"1"`
	Title       string `json:"title" example:"Buy groceries"`
	Description string `json:"description" example:"Buy milk, bread, and cheese"`
}
