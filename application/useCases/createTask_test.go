package usecases_test

import (
	usecases "jeanmrtns/sample-go-api/application/useCases"
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/infra/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTaskWithEmptyTitle(t *testing.T) {
	r := repositories.InMemoryTaskRepository{}
	err := usecases.CreateTask(&r, &domain.Task{
		Title:       "",
		Description: "Some description",
		Done:        false,
	})

	assert.NotEmpty(t, err)
}

func TestCreateTaskWithValidTitle(t *testing.T) {
	r := repositories.InMemoryTaskRepository{}
	err := usecases.CreateTask(&r, &domain.Task{
		Title:       "Some title",
		Description: "Some description",
		Done:        false,
	})

	assert.Empty(t, err)
}

func TestAllCreatedTasksMustNotBeDone(t *testing.T) {
	r := repositories.InMemoryTaskRepository{}

	usecases.CreateTask(&r, &domain.Task{
		Title:       "Some title",
		Description: "Some description",
		Done:        true,
	})

	usecases.CreateTask(&r, &domain.Task{
		Title:       "Some title 2",
		Description: "Some description 2",
		Done:        false,
	})

	tasks := usecases.GetAllTasks(&r)

	assert.Len(t, tasks, 2)
	assert.False(t, tasks[0].Done)
	assert.False(t, tasks[1].Done)
}
