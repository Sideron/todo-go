package models

type Task struct {
	name        string
	description string
	completed   bool
}

func NewTask(name string, description string) *Task {
	return &Task{
		name:        name,
		description: description,
		completed:   false,
	}
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) IsCompleted() bool {
	return t.completed
}
