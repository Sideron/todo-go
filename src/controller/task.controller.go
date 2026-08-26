package controller

import "todo-go/src/models"

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

func (c *TaskController) Tasks() []models.Task {
	nTasks := make([]models.Task, 0, len(c.tasks))
	for i := range len(c.tasks) {
		nTasks = append(nTasks, *c.tasks[i])
	}
	return nTasks
}

func (c *TaskController) CreateNewTask(name string, description string) {
	for i := range len(c.tasks) {
		if c.tasks[i].Name() == name {
			return
		}
	}
	c.tasks = append(c.tasks, models.NewTask(name, description))
}
