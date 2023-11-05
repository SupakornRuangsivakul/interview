package interview

import (
	"context"
	"interview-rbh/model"
	"interview-rbh/service/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cardResponse struct {
	ID         string `json:"id"`
	CardStatus string `json:"cardStatus" `
	CardDetail string `json:"cardDetail" `
	CardName   string `json:"cardName" `
	Email      string `json:"email" `
	CreatedBy  string `json:"createBy" `
	CreatedAt  string `json:"createAt" `
}

func FetchCard(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := model.GetMongoClient(ctx).Database("interviewrbh").Collection("interviewCard")
	filter := bson.M{"cardStatus": bson.M{"$ne": StatusKeep}}

	// Find multiple documents
	var cardLists []interviewCard // A slice to hold the query results
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "createAt", Value: -1}})
	cur, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	defer cur.Close(context.Background())

	// Iterate through the cursor
	for cur.Next(context.Background()) {
		var elem interviewCard
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		cardLists = append(cardLists, elem)
	}

	if err := cur.Err(); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	results := buildCardResponse(cardLists)
	c.JSON(200, results)

}

func buildCardResponse(cards []interviewCard) []cardResponse {
	var result []cardResponse
	for _, card := range cards {
		item := cardResponse{
			ID:         card.ID.Hex(),
			CardStatus: card.CardStatus,
			CardDetail: card.CardDetail,
			CardName:   card.CardName,
			Email:      card.Email,
			CreatedBy:  card.CreatedBy,
			CreatedAt:  utils.GetThaiDate(card.CreatedAt),
		}
		result = append(result, item)
	}
	return result
}
