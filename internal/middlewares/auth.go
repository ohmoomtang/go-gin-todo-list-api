package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"oot.me/todo-list-api/utils"
)


func ValidateAuthenticationToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
    if authHeader == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        ctx.Abort()
        return
    }
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    if tokenString == authHeader {
        ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        ctx.Abort()
        return
    }
	claims,err := utils.ValidateJWT(tokenString)
	if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        ctx.Abort()
        return
    }
	ctx.Set("uid",claims["uid"])
	ctx.Set("name",claims["name"])
	ctx.Set("email",claims["email"])
	ctx.Set("exp",claims["exp"])
	ctx.Next()
}