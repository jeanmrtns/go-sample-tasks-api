package repositories

import (
	"jeanmrtns/sample-go-api/domain"
	"time"
)

type InMemoryTaskRepository struct {
	Tasks []domain.Task
}

func (s *InMemoryTaskRepository) FindAll() []domain.Task {
	return s.Tasks
}

func (s *InMemoryTaskRepository) Create(task *domain.Task) error {
	s.Tasks = append(s.Tasks, domain.Task{
		Title:       task.Title,
		Description: task.Description,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	return nil
}
