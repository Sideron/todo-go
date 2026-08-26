package controller

import "todo-go/src/models"

type TaskController struct {
	tasks []models.Task
}

func (c *TaskController) completeTask(name string) {
	for i := range len(c.tasks) {
		if c.tasks[i].Name() == name {
			c.tasks[i].Complete()
			return
		}
	}
}

func (c *TaskController) Tasks() []models.Task {
	return c.tasks
}

func (c *TaskController) createNewTask(name string, description string) {
	for i := range len(c.tasks) {
		if c.tasks[i].Name() == name {
			return
		}
	}
	nTask := models.NewTask(name, description)
	c.tasks = append(c.tasks, *nTask)
}
