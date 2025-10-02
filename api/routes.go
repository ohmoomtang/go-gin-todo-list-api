package api

import (
	"github.com/gin-gonic/gin"
	"oot.me/todo-list-api/internal/handlers"
)

func RegisterRoute(server *gin.Engine) {
	server.POST("/register",handlers.RegisterUser)
	server.POST("/login",handlers.Login)
	server.POST("/todos",handlers.CreateTodo)
	server.PUT("/todos/:id",handlers.UpdateTodo)
	server.DELETE("/todos/:id",handlers.DeleteTodo)
	server.GET("/todos",handlers.FetchTodo)
}