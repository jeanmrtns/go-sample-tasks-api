package handler

import (
	usecases "jeanmrtns/sample-go-api/application/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasksHandler(c *gin.Context) {
	tasks, err := usecases.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error taking tasks",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    tasks,
	})
}
