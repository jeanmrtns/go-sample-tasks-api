package usecases

import (
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/infra/repositories"
)

func GetAllTasks(r repositories.TaskRepository) []domain.Task {
	tasks := r.FindAll()

	return tasks
}
