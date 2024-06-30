package repositories

import (
	"jeanmrtns/sample-go-api/config"
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/models"
)

type SQLiteTaskRepository struct{}

func (s SQLiteTaskRepository) FindAll() []domain.Task {
	var tasks []models.Task
	config.DB.Find(&tasks)

	var parsedTasks []domain.Task
	for _, task := range tasks {
		parsedTasks = append(parsedTasks, domain.Task{
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return parsedTasks
}

func (s SQLiteTaskRepository) Create(task *domain.Task) {
	config.DB.Create(&models.Task{
		Title:       task.Title,
		Description: task.Description,
		Done:        false,
	})
}
