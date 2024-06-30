package router

import (
	"jeanmrtns/sample-go-api/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/tasks", handler.ListTasksHandler)
		v1.GET("/task", handler.GetTaskHandler)
		v1.POST("/task", handler.CreateTaskHandler)
		v1.PUT("/task", handler.UpdateTaskHandler)
		v1.DELETE("/task", handler.DeleteTaskHandler)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
