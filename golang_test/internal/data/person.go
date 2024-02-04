package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://127.0.0.1:27017/"//+"?maxPoolSize=20&w=majority"
// const uri = "mongodb://user:admin@127.0.0.1:27017/"+"?maxPoolSize=20&w=majority"

// type Person struct {
// 	Id   int    `bson:"id"`
// 	Name string `bson:name`
// }

func GetAdminDatabase() *mongo.Database {
	
	var result bson.M //宣告result為bson.M型別
	client := getDatabaseClient() //宣告client+指派內容
	database := client.Database("admin")

	// Send a ping to confirm a successful connection
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return database

}

func getDatabaseClient() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	

	return client
}