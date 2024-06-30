package usecases

import (
	"errors"
	"fmt"
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/infra/repositories"
)

func CreateTask(r repositories.TaskRepository, task *domain.Task) error {
	if task.Title == "" {
		return errors.New("title should not be empty")
	}

	err := r.Create(&domain.Task{
		Title:       task.Title,
		Description: task.Description,
		Done:        false,
	})

	if err != nil {
		fmt.Printf("Error creating task %v", err)
		return err
	}

	return nil
}
