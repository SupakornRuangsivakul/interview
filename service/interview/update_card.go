package interview

import (
	"context"
	"fmt"
	"interview-rbh/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type updateCard struct {
	CardID     string `json:"cardId"`
	CardName   string `json:"cardName"`
	CardDetail string `json:"cardDetail"`
	CardStatus string `json:"cardStatus"`
}

func UpdateCard(c *gin.Context) {
	user := c.GetString("userId")
	fmt.Println(user)
	var req updateCard
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Println(req)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)

	// session, err := client.StartSession()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer session.EndSession(context.Background())

	// Define the transaction options
	// wc := writeconcern.Majority()
	// transactionOptions := options.Transaction().
	// 	SetReadConcern(readconcern.Local()).
	// 	SetWriteConcern(wc)

	// // Define the work to be done in the transaction
	// transaction := func(sessCtx mongo.SessionContext) (interface{}, error) {
	// Perform operations within the transaction
	// client := model.GetMongoClient(ctx)
	card, err := getCardByID(req.CardID, client, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Println(card)
	collection := client.Database("interviewrbh").Collection("changeLogCard")
	now := time.Now()
	changeLogCard := changeLogCard{
		CardID:     card.ID.Hex(),
		CardDetail: card.CardDetail,
		CardName:   card.CardName,
		CardStatus: card.CardStatus,
		CreatedBy:  user,
		CreatedAt:  now,
	}
	_, err = collection.InsertOne(context.Background(), changeLogCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	cardUpdate := bson.M{
		"$set": bson.M{
			"cardDetail": req.CardDetail,
			"cardStatus": req.CardStatus,
			"cardName":   req.CardName,
			"updateAt":   time.Now(),
		},
	}
	collection = client.Database("interviewrbh").Collection("interviewCard")
	updateSuccess, err := collection.UpdateByID(context.Background(), card.ID, cardUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, updateSuccess)
	// }
	// // Execute the transaction
	// _, err = session.WithTransaction(context.Background(), transaction, transactionOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Transaction committed successfully")
	// }

}
