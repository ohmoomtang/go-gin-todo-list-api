package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"oot.me/todo-list-api/internal/models"
)

func CreateTodo(ctx *gin.Context){
	uid,err := strconv.ParseFloat(ctx.GetString("uid"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	exp,err := strconv.ParseFloat(ctx.GetString("exp"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if uid == 0 || int64(exp) == 0 || time.Now().After(time.Unix(int64(exp), 0)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	var newTodo models.Todo
	err = ctx.ShouldBindJSON(&newTodo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newTodo.UserId = int64(uid)
	err = newTodo.New()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{"id": newTodo.ID,"title": newTodo.Title,"description": newTodo.Description})
}

func UpdateTodo(ctx *gin.Context){
	uid,err := strconv.ParseFloat(ctx.GetString("uid"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	exp,err := strconv.ParseFloat(ctx.GetString("exp"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if uid == 0 || int64(exp) == 0 || time.Now().After(time.Unix(int64(exp), 0)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
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
	if existingTodo.UserId != int64(uid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
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
	uid,err := strconv.ParseFloat(ctx.GetString("uid"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	exp,err := strconv.ParseFloat(ctx.GetString("exp"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if uid == 0 || int64(exp) == 0 || time.Now().After(time.Unix(int64(exp), 0)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	deletedTodoId := ctx.Param("id")
	deletedTodoIdConv,err := strconv.ParseInt(deletedTodoId,10,64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot parsed ID"})		
		return
	}
	deletedTodo,err := models.FindSingleTodo(deletedTodoIdConv)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})		
		return
	}
	if deletedTodo.UserId != int64(uid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err = models.Delete(deletedTodoIdConv)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func FetchSingleTodo(ctx *gin.Context){
	uid,err := strconv.ParseFloat(ctx.GetString("uid"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	exp,err := strconv.ParseFloat(ctx.GetString("exp"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if uid == 0 || int64(exp) == 0 || time.Now().After(time.Unix(int64(exp), 0)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
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
	if todo.UserId != int64(uid) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"id": todo.ID,"title": todo.Title,"description": todo.Description})
}

func FetchTodo(ctx *gin.Context){
	uid,err := strconv.ParseFloat(ctx.GetString("uid"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	exp,err := strconv.ParseFloat(ctx.GetString("exp"),64)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if uid == 0 || int64(exp) == 0 || time.Now().After(time.Unix(int64(exp), 0)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	var limitQueryConv int64 = 10
	var pageQueryConv int64 = 1
	var todoList []models.Todo
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
	todoList,err = models.FetchTodos(int64(uid),pageQueryConv,limitQueryConv)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})		
		return
	}
	if len(todoList) == 0 {
		todoList = make([]models.Todo,0)
	}
	ctx.JSON(http.StatusOK,gin.H{"data": todoList,"page": pageQueryConv,"limit": limitQueryConv,"total":len(todoList)})
}