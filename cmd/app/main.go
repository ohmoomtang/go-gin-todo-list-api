package main

import (
	"github.com/gin-gonic/gin"
	"oot.me/todo-list-api/api"
)

func main (){
	server := gin.Default()
	api.RegisterRoute(server)
	server.Run()
}