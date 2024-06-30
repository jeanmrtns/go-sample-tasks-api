package repositories

import (
	"jeanmrtns/sample-go-api/domain"
)

type TaskRepository interface {
	FindAll() []domain.Task
	Create(task *domain.Task) error
}
