package interview

import (
	"context"
	"interview-rbh/model"
	"interview-rbh/service/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type changeLogCard struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CardID     string             `bson:"cardId"`
	CardName   string             `bson:"cardName" json:"cardName"`
	CardDetail string             `bson:"cardDetail" json:"cardDetail"`
	CardStatus string             `bson:"cardStatus" json:"cardStatus"`
	CreatedBy  string             `bson:"createBy" json:"createBy"`
	CreatedAt  time.Time          `bson:"createAt" json:"createAt"`
}

type changeLogCardResposne struct {
	CardID     string `json:"cardId"`
	CardName   string `json:"cardName"`
	CardDetail string `json:"cardDetail"`
	CardStatus string `json:"cardStatus"`
	CreatedBy  string `json:"createBy"`
	CreatedAt  string `json:"createAt"`
}

func GetHistoryByCardID(c *gin.Context) {
	idStr := c.Param("cardId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)

	collection := client.Database("interviewrbh").Collection("changeLogCard")
	filter := bson.M{"cardId": idStr}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "createAt", Value: -1}})
	cur, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(context.Background())
	var changeLogCards []changeLogCard
	for cur.Next(context.Background()) {
		var elem changeLogCard
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		changeLogCards = append(changeLogCards, elem)
	}
	result := buildResponse(changeLogCards)
	c.JSON(200, result)
}

func buildResponse(changeLogCards []changeLogCard) []changeLogCardResposne {
	var result []changeLogCardResposne
	for _, changeLogCard := range changeLogCards {
		item := changeLogCardResposne{
			CardID:     changeLogCard.CardID,
			CardName:   changeLogCard.CardName,
			CardDetail: changeLogCard.CardDetail,
			CardStatus: changeLogCard.CardStatus,
			CreatedBy:  changeLogCard.CreatedBy,
			CreatedAt:  utils.GetThaiDate(changeLogCard.CreatedAt),
		}
		result = append(result, item)
	}
	return result
}
