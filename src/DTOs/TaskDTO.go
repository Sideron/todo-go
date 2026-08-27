package dtos

import "todo-go/src/models"

type TaskDTO struct {
	Name        string
	Description string
	Completed   bool
}

func TaskToDTO(t models.Task) TaskDTO {
	return TaskDTO{
		Name:        t.Name(),
		Description: t.Description(),
		Completed:   t.IsCompleted(),
	}
}
