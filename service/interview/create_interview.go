package interview

import (
	"context"
	"interview-rbh/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	StatusTodo       = "To Do"
	StatusInProgress = "In Progress"
	StatusDone       = "Done"
	StatusKeep       = "Keep"
)

type interviewCard struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CardStatus string             `bson:"cardStatus" json:"cardStatus"`
	CardDetail string             `bson:"cardDetail" json:"cardDetail"`
	CardName   string             `bson:"cardName" json:"cardName"`
	Email      string             `bson:"email" json:"email"`
	CreatedBy  string             `bson:"createBy" json:"createBy"`
	CreatedAt  time.Time          `bson:"createAt" json:"createAt"`
	UpdatedAt  time.Time          `bson:"updateAt" json:"updateAt"`
}

func CreateInterview() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)
	database := client.Database("interviewrbh")

	// Create a collection
	collectionName := "interviewCard"
	if err := database.CreateCollection(ctx, collectionName); err != nil {
		log.Fatal(err)
	}
	now := time.Now()

	cardList := []interviewCard{
		{CardStatus: "To Do", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 1", Email: "User01@robinhood.co.th", CreatedBy: "User01", CreatedAt: now, UpdatedAt: now},
		{CardStatus: "To Do", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 2", Email: "User01@robinhood.co.th", CreatedBy: "User01", CreatedAt: now.Add(time.Hour * 2), UpdatedAt: now.Add(time.Hour * 2)},
		{CardStatus: "To Do", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 3", Email: "User01@robinhood.co.th", CreatedBy: "User01", CreatedAt: now.Add(time.Hour * 3), UpdatedAt: now.Add(time.Hour * 3)},
		{CardStatus: "In Progress", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 4", Email: "User02@robinhood.co.th", CreatedBy: "User02", CreatedAt: now.Add(time.Hour), UpdatedAt: now.Add(time.Hour)},
		{CardStatus: "To Do", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 5", Email: "User02@robinhood.co.th", CreatedBy: "User02", CreatedAt: now.Add(time.Hour * 4), UpdatedAt: now.Add(time.Hour * 4)},
		{CardStatus: "To Do", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 6", Email: "User02@robinhood.co.th", CreatedBy: "User02", CreatedAt: now.Add(time.Hour * 6), UpdatedAt: now.Add(time.Hour * 6)},
		{CardStatus: "In Progress", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 7", Email: "User03@robinhood.co.th", CreatedBy: "User03", CreatedAt: now.Add(time.Hour * 5), UpdatedAt: now.Add(time.Hour * 5)},
		{CardStatus: "To Do", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 8", Email: "User03@robinhood.co.th", CreatedBy: "User03", CreatedAt: now.Add(time.Hour * 7), UpdatedAt: now.Add(time.Hour * 7)},
		{CardStatus: "Done", CardDetail: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", CardName: "นัดสัมภาษณ์งาน 9", Email: "User03@robinhood.co.th", CreatedBy: "User03", CreatedAt: now.Add(time.Hour * 8), UpdatedAt: now.Add(time.Hour * 8)},
	}
	var documents []interface{}
	for _, card := range cardList {
		documents = append(documents, card)
	}
	collection := client.Database("interviewrbh").Collection("interviewCard")
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted documents with IDs: %v\n", result.InsertedIDs)

	if err := database.CreateCollection(ctx, "commentCard", nil); err != nil {
		log.Fatal(err)
	}

	if err := database.CreateCollection(ctx, "changeLogCard", nil); err != nil {
		log.Fatal(err)
	}

}
