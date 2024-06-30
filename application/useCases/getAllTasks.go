package usecases

import (
	"jeanmrtns/sample-go-api/config"
	"jeanmrtns/sample-go-api/models"
)

func GetAllTasks() ([]models.TaskResponse, error) {
	var tasks []models.Task
	result := config.DB.Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	var parsedTasks []models.TaskResponse
	for _, task := range tasks {
		parsedTasks = append(parsedTasks, models.TaskResponse{
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return parsedTasks, nil
}
