package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Database *mongo.Database

const BooksCollection = "books"

func Connect() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("database connection error ", err)
		panic(err)
	}
	return client.Database("test_db")
}
