package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"oot.me/todo-list-api/pkg/utils"
)

type Todo struct {
	ID bson.ObjectID
	Title string `json:"title" bingind:"required"`
	Description string `json:"description" bingind:"required"`
	UserId int64
}

func (t *Todo) New() (error){
	collection := utils.MongoDB.Collection("todo")
	result,err := collection.InsertOne(context.TODO(),t)
	if err != nil {
		return err
	}
	t.ID = result.InsertedID.(bson.ObjectID)
	return nil
}