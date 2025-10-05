package main

import (
	"github.com/gin-gonic/gin"
	"oot.me/todo-list-api/api"
	"oot.me/todo-list-api/config"
	"oot.me/todo-list-api/pkg/utils"
)

func main (){
	utils.ConnectMongoDB(config.MONGODB_URI)
	server := gin.Default()
	api.RegisterRoute(server)
	server.Run()
}