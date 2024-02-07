package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const uri = "mongodb://mongodb:27017/"+"?maxPoolSize=20&w=majority"
var client *mongo.Client
var err error

// get admin database connection
func GetAdminDatabase() *mongo.Database {
	log.SetPrefix("")
	
	// var result bson.M //宣告result為bson.M型別
	client = getDatabaseClient() //宣告client+指派內容
	database := client.Database("admin")

	var result bson.M
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	disconnectDatabaseClient()
	return database

}

// get go project database connection
func getDatabase() *mongo.Database {
	log.SetPrefix("")
	
	// var result bson.M //宣告result為bson.M型別
	client = getDatabaseClient() //宣告client+指派內容
	database := client.Database("golangdb")

	var result bson.M
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to golangdb!")
	return database

}

// get mongo connection
func getDatabaseClient() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// var result bson.M
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func disconnectDatabaseClient() {
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func InsData(tableName, data) {
	getDatabase()
	coll := database.Collection(tableName)
	result, err := client.Database("admin").coll.InsertOne(context.TODO(), data)
	processEDataError(err)
	disconnectDatabaseClient()
	return result, err
}

func InsMulData(tableName, data) {
	getDatabase()
	coll := database.Collection(tableName)
	result, err := client.Database("admin").InsertMany(context.TODO(), data)
	processEDataError(err)
	disconnectDatabaseClient()
	return result, err

}

func FindData(tableName, data) {
	getDatabase()
	coll := database.Collection(tableName)
	cursor, err := coll.Find(context.TODO(), data)
	err = coll.FindOne(context.TODO(), data).Decode(&result)
	processEDataError(err)
	disconnectDatabaseClient()
	return cursor, err
}

func FindMulData(tableName, data) {
	getDatabase()
	coll := database.Collection(tableName)
	cursor, err := coll.Find(context.TODO(), data)
	processEDataError(err)
	disconnectDatabaseClient()
	return cursor, err
}

func processEDataError(err) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
}