package controller

import (
	dtos "todo-go/src/DTOs"
	"todo-go/src/models"
)

type TaskController struct {
	tasks []*models.Task
}

func (c *TaskController) CompleteTask(name string) {
	for i := range len(c.tasks) {
		if c.tasks[i].Name() == name {
			c.tasks[i].Complete()
			return
		}
	}
}

func (c *TaskController) Tasks() []dtos.TaskDTO {
	nTasks := make([]dtos.TaskDTO, 0, len(c.tasks))
	for _, t := range c.tasks {
		nTasks = append(nTasks, dtos.TaskToDTO(*t))
	}
	return nTasks
}

func (c *TaskController) CreateNewTask(name string, description string) {
	for _, t := range c.tasks {
		if t.Name() == name {
			return
		}
	}
	c.tasks = append(c.tasks, models.NewTask(name, description))
}
