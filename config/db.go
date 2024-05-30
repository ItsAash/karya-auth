// config/db.go
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Replace the placeholder with your Atlas connection string
const uri = "mongodb+srv://aash:root@karya.sws8xul.mongodb.net/?retryWrites=true&w=majority&appName=karya"

func Connect() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Ping the database to ensure connection is established
	err = pingDB(Client)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB")
}

func pingDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return err
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	databaseName := "karya-test" // Replace with your actual database name
	return Client.Database(databaseName).Collection(collectionName)
}