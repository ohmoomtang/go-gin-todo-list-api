package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"oot.me/todo-list-api/config"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database
var err error

func ConnectMongoDB(uri string){
	if uri == "" {
		log.Fatal("No MongoDB URI found. You must set your 'MONGODB_URI' environment variable.")
	}
	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Creates a new client and connects to the server
	MongoClient, err = mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := MongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	MongoDB = MongoClient.Database(config.MONGODB_DB_NAME)
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}