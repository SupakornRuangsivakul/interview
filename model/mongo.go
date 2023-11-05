package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetMongoClient(ctx context.Context) *mongo.Client {

	if err := client.Ping(ctx, nil); err != nil {
		connectMongoDB(ctx)
	}
	return client
}

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connectMongoDB(ctx)
}

func connectMongoDB(ctx context.Context) {
	uri := "mongodb://root:password@mongo:27017"
	// Create a new client and connect to the server
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	// Call Ping to verify that the connection is alive
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected and pinged.")
}
