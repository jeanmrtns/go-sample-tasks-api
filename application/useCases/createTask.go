package usecases

import (
	"jeanmrtns/sample-go-api/config"
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/models"
)

func CreateTask(task *domain.Task) {
	config.DB.Create(&models.Task{
		Title:       task.Title,
		Description: task.Description,
		Done:        false,
	})
}
