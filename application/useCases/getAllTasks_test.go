package usecases_test

import (
	usecases "jeanmrtns/sample-go-api/application/useCases"
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/infra/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmptyListWhenNoTaskCreated(t *testing.T) {
	r := repositories.InMemoryTaskRepository{}
	tasks := usecases.GetAllTasks(&r)

	assert.Len(t, tasks, 0)
	assert.Empty(t, tasks)
}

func TestGetPopulatedListWhenAreTasksCreated(t *testing.T) {
	r := repositories.InMemoryTaskRepository{}

	usecases.CreateTask(&r, &domain.Task{
		Title:       "Some title",
		Description: "Some description",
		Done:        false,
	})

	usecases.CreateTask(&r, &domain.Task{
		Title:       "Some title 2",
		Description: "Some description 2",
		Done:        false,
	})

	tasks := usecases.GetAllTasks(&r)

	assert.Len(t, tasks, 2)

	assert.NotEmpty(t, tasks)
	assert.Equal(t, "Some title", tasks[0].Title)
	assert.Equal(t, "Some description", tasks[0].Description)
	assert.False(t, tasks[0].Done)

	assert.Equal(t, "Some title 2", tasks[1].Title)
	assert.Equal(t, "Some description 2", tasks[1].Description)
	assert.False(t, tasks[1].Done)
}
