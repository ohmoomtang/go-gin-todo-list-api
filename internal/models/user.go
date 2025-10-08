package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"oot.me/todo-list-api/utils"
)

type User struct {
	_id bson.ObjectID `json:"_id"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *User) New() (error) {
	opts := options.Count().SetHint("_id_")
	collection := utils.MongoDB.Collection("users")
	docCount,err := collection.CountDocuments(context.TODO(),bson.D{},opts)
	if err != nil {
		return err
	}
	u.ID = docCount + 1
	_,err = collection.InsertOne(context.TODO(),u)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(email string) (User,error) {
	collection := utils.MongoDB.Collection("users")
	filter := bson.D{{"email",email},}
	result := collection.FindOne(context.TODO(),filter)
	if result.Err() != nil {
		return User{},result.Err()
	}
	var resultUser User
	err := result.Decode(&resultUser)
	if err != nil {
		return User{},err
	}
	return resultUser,nil
}