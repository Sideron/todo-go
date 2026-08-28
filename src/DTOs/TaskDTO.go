package dtos

import "todo-go/src/models"

type TaskDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func TaskToDTO(t models.Task) TaskDTO {
	return TaskDTO{
		Name:        t.Name(),
		Description: t.Description(),
		Completed:   t.IsCompleted(),
	}
}
