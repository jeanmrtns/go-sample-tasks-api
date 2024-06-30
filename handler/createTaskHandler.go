package handler

import (
	usecases "jeanmrtns/sample-go-api/application/useCases"
	"jeanmrtns/sample-go-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	task := domain.NewTask(title, description, false)

	usecases.CreateTask(&task)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
