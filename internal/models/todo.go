package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"oot.me/todo-list-api/utils"
)

type Todo struct {
	_id bson.ObjectID `json:"_id"`
	ID int64 `json:"id"`
	Title string `json:"title" bingind:"required"`
	Description string `json:"description" bingind:"required"`
	UserId int64 `json:"userid"`
}

func (t *Todo) New() (error){
	opts := options.Count().SetHint("_id_")
	collection := utils.MongoDB.Collection("todo")
	docCount,err := collection.CountDocuments(context.TODO(),bson.D{},opts)
	if err != nil {
		return err
	}
	t.ID = docCount + 1 
	_,err = collection.InsertOne(context.TODO(),t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) Update() (error){
	collection := utils.MongoDB.Collection("todo")
	filter := bson.D{{"id",t.ID}}
	_,err := collection.ReplaceOne(context.TODO(), filter, t)
	if err != nil {
		return err
	}
	return nil
}

func FetchTodos(uid,page,limit int64) ([]Todo,error) {
	collection := utils.MongoDB.Collection("todo")
	filter := bson.D{{"userid",uid}}

	skip := int64((page - 1) * int64(limit))

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	
	cursor,err := collection.Find(context.TODO(),filter,findOptions)
	if err != nil {
		return []Todo{},err
	}
	defer cursor.Close(context.TODO())
	var todoList []Todo
	err = cursor.All(context.TODO(),&todoList)
	if err != nil {
		return []Todo{},err
	}
	return todoList,nil
}

func FindSingleTodo(id int64) (Todo, error) {
	collection := utils.MongoDB.Collection("todo")
	filter := bson.D{{"id",id}}
	result := collection.FindOne(context.TODO(),filter)
	if result.Err() != nil {
		return Todo{},result.Err()
	}
	var resultTodo Todo
	err := result.Decode(&resultTodo)
	if err != nil {
		return Todo{},err
	}
	return resultTodo,nil
}

func Delete(id int64) (error){
	collection := utils.MongoDB.Collection("todo")
	filter := bson.D{{"id",id}}
	_,err := collection.DeleteOne(context.TODO(),filter)
	if err != nil {
		return err
	}
	return nil
}