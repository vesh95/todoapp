package todo

import "github.com/google/uuid"

// Todo Структура
type Todo struct {
	ID         uuid.UUID `json:"id"`
	Task       string    `json:"task"`
	IsComplete bool      `json:"is_complete"`
}
