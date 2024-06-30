package handler

import (
	usecases "jeanmrtns/sample-go-api/application/useCases"
	"jeanmrtns/sample-go-api/infra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasksHandler(c *gin.Context) {
	r := repositories.SQLiteTaskRepository{}
	tasks := usecases.GetAllTasks(r)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    tasks,
	})
}
