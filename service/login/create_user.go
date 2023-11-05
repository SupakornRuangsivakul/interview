package login

import (
	"context"
	"interview-rbh/model"
	"log"
	"time"
)

type User struct {
	UserName  string    `bson:"userName"`
	Password  string    `bson:"password"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Role      string    `bson:"role"`
	Level     string    `bson:"level"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func CreateUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)
	database := client.Database("interviewrbh")

	// opts := options.CreateCollection().SetMaxDocuments(1024 * 1024)

	// Create a collection
	collectionName := "UserRole"
	if err := database.CreateCollection(ctx, collectionName, nil); err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	userList := []User{
		{UserName: "User01", Password: hashPassword("password"), Name: "robinhood1", Email: "User1@robinhood.co.th", Level: "1", Role: "role1", CreatedAt: now, UpdatedAt: now},
		{UserName: "User02", Password: hashPassword("password"), Name: "robinhood2", Email: "User2@robinhood.co.th", Level: "2", Role: "role2", CreatedAt: now, UpdatedAt: now},
		{UserName: "User03", Password: hashPassword("password"), Name: "robinhood3", Email: "User3@robinhood.co.th", Level: "3", Role: "role3", CreatedAt: now, UpdatedAt: now},
	}
	var documents []interface{}
	for _, user := range userList {
		documents = append(documents, user)
	}
	collection := client.Database("interviewrbh").Collection("UserRole")
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted documents with IDs: %v\n", result.InsertedIDs)
}
