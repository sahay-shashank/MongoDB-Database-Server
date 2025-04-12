package handler

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDatabase() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("URI not set in 'MONGODB_URI' environment variable.")
		return nil, fmt.Errorf("the URI is not set in 'MONGODB_URI' environment variable.")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB URI.")
		return nil, fmt.Errorf("error connecting to MongoDB. Error: %v", err)
	}
	log.Printf("Connected to MongoDB")
	return client, nil
}
