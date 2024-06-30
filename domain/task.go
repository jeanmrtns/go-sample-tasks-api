package domain

type Task struct {
	Title       string
	Description string
	Done        bool
}

func NewTask(title, description string, done bool) Task {
	task := Task{
		Title:       title,
		Description: description,
		Done:        done,
	}

	return task
}
