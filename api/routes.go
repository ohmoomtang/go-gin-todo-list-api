package api

import (
	"github.com/gin-gonic/gin"
	"oot.me/todo-list-api/internal/handlers"
	"oot.me/todo-list-api/internal/middlewares"
)

func RegisterRoute(server *gin.Engine) {
	server.POST("/register",handlers.RegisterUser)
	server.POST("/login",handlers.Login)
	server.POST("/todos",middlewares.ValidateAuthenticationToken,handlers.CreateTodo)
	server.PUT("/todos/:id",middlewares.ValidateAuthenticationToken,handlers.UpdateTodo)
	server.DELETE("/todos/:id",middlewares.ValidateAuthenticationToken,handlers.DeleteTodo)
	server.GET("/todos/:id",middlewares.ValidateAuthenticationToken,handlers.FetchSingleTodo)
	server.GET("/todos",middlewares.ValidateAuthenticationToken,handlers.FetchTodo)
}