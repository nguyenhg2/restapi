package routes

import (
	"github.com/gin-gonic/gin"

	"restapi/handlers"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", handlers.GetTodos)
		v1.POST("/todos", handlers.CreateTodo)
		v1.DELETE("/todos/:id", handlers.DeleteTodo)
	}
}
