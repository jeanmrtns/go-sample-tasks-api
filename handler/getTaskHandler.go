package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "One task",
	})
}
