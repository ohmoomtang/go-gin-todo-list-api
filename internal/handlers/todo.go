package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	updatedTodoId := ctx.Param("id")
	updatedTodoIdConv,err := strconv.ParseInt(updatedTodoId,10,64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot parsed ID"})		
		return
	}
	var existingTodo models.Todo
	existingTodo,err = models.FindSingleTodo(updatedTodoIdConv)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//TODO : logic to retrieve current userId, verify, etc.
	var updatedTodo models.Todo
	err = ctx.ShouldBindJSON(&updatedTodo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	existingTodo.Title = updatedTodo.Title
	existingTodo.Description = updatedTodo.Description
	err = existingTodo.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"id": existingTodo.ID,"title": existingTodo.Title,"description": existingTodo.Description})
}

func DeleteTodo(ctx *gin.Context){
	deletedTodoId := ctx.Param("id")
	deletedTodoIdConv,err := strconv.ParseInt(deletedTodoId,10,64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot parsed ID"})		
		return
	}
	//TODO : logic to retrieve current userId, verify, etc.
	err = models.Delete(deletedTodoIdConv)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func FetchSingleTodo(ctx *gin.Context){
	fetchedTodoId := ctx.Param("id")
	fetchedTodoIdConv,err := strconv.ParseInt(fetchedTodoId,10,64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot parsed ID"})		
		return
	}
	var todo models.Todo
	todo,err = models.FindSingleTodo(fetchedTodoIdConv)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//TODO : logic to retrieve current userId, verify, etc. - if not associated user then cannot return resources
	ctx.JSON(http.StatusOK,gin.H{"id": todo.ID,"title": todo.Title,"description": todo.Description})
}

func FetchTodo(ctx *gin.Context){
	var limitQueryConv int64 = 10
	var pageQueryConv int64 = 1
	var todoList []models.Todo
	var err error
	pageQuery := ctx.Query("page")
	limitQuery := ctx.Query("limit")
	if pageQuery != "" {
		pageQueryConv,err = strconv.ParseInt(pageQuery,10,64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot parsed page query params"})		
			return
		}		
	}
	if limitQuery != "" {
		limitQueryConv,err = strconv.ParseInt(limitQuery,10,64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot parsed limit query params"})		
			return
		}		
	}
	//Hardcode UID first
	todoList,err = models.FetchTodos(1,pageQueryConv,limitQueryConv)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})		
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"data": todoList,"page": pageQueryConv,"limit": limitQueryConv,"total":len(todoList)})
}