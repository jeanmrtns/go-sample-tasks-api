package domain

import "time"

type Task struct {
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(title, description string, done bool) Task {
	task := Task{
		Title:       title,
		Description: description,
		Done:        done,
	}

	return task
}
