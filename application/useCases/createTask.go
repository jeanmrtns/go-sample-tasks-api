package usecases

import (
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/infra/repositories"
)

func CreateTask(r repositories.TaskRepository, task *domain.Task) {
	r.Create(&domain.Task{
		Title:       task.Title,
		Description: task.Description,
		Done:        false,
	})
}
