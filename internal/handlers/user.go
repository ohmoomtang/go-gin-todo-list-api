package handlers

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"oot.me/todo-list-api/internal/models"
	"oot.me/todo-list-api/utils"
)

func RegisterUser(ctx *gin.Context){
	var newUser models.User
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = mail.ParseAddress(newUser.Email)
    if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email was incorrect format"})
		return
    }
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(newUser.Password),bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	newUser.Password = string(hashedPassword)
	err = newUser.New()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	jwtToken,err := utils.SignJWT(newUser.ID,newUser.Name,newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{"token": jwtToken})
}

func Login(ctx *gin.Context){
	var loginUser models.User
	err := ctx.ShouldBindJSON(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = mail.ParseAddress(loginUser.Email)
    if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email was incorrect format"})
		return
    }
	foundUser,err := models.FindUser(loginUser.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Email or Password"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password),[]byte(loginUser.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Email or Password"})
		return
	}
	token,err := utils.SignJWT(foundUser.ID,foundUser.Name,foundUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"token": token})
}