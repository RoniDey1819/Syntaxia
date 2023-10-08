package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "syntaxia"
const attendeesColName = "attendees"

var attendeesCollection *mongo.Collection


var client *mongo.Client

func CreateDB(mongoUri string) (*mongo.Collection, error) {
	clientOption := options.Client().ApplyURI(mongoUri)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("MongoDB connection success")
	attendeesCollection = client.Database(dbName).Collection(attendeesColName)
	fmt.Println("Collection instance is ready")

	return attendeesCollection, nil	
}



func CloseDB() {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("MongoDB connection closed")
	}
}