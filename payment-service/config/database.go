package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	db := os.Getenv("MONGO_DB")
	if db == "" {
		log.Fatal("MONGO_DB environment variable is not set")
	}
	collection := client.Database(db).Collection(collectionName)
	return collection
}
