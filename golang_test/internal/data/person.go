package data

import (
	// "fmt"
	"log"
)


type Person struct {
	Id   int    `bson:"id"`
	Name string `bson:name`
	FirstName string `bson:firstname`
	LastName string `bson:lastname`
	Age int `int:age`
}

func CreatePerson() {
	docs := []interface{} {
	    bson.D{{"firstName", "Erik"}, {"age", 27}},
	    bson.D{{"firstName", "Mohammad"}, {"lastName", "Ahmad"}, {"age", 10}},
	    bson.D{{"firstName", "Todd"}},
	    bson.D{{"firstName", "Juan"}, {"lastName", "Pablo"}}
	}


	result, err := InsMulData()

	log.Println(result)
	log.Println(err)

	// defer disconnectDatabaseClient()
	// return database

}