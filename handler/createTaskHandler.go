package handler

import (
	usecases "jeanmrtns/sample-go-api/application/useCases"
	"jeanmrtns/sample-go-api/domain"
	"jeanmrtns/sample-go-api/infra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	task := domain.NewTask(title, description, false)

	r := repositories.SQLiteTaskRepository{}

	usecases.CreateTask(&r, &task)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
