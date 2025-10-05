package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"oot.me/todo-list-api/internal/models"
)

func CreateTodo(ctx *gin.Context){
	var newTodo models.Todo
	err := ctx.ShouldBindJSON(&newTodo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//TODO : logic to retrieve current userId and update to newTodo object
	newTodo.UserId = 1
	err = newTodo.New()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{"id": newTodo.ID,"title": newTodo.Title,"description": newTodo.Description})
}

func UpdateTodo(ctx *gin.Context){

}

func DeleteTodo(ctx *gin.Context){

}

func FetchTodo(ctx *gin.Context){

}