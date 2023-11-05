package interview

import (
	"context"
	"interview-rbh/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func KeepInterviewCard(c *gin.Context) {
	idStr := c.Param("cardId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)

	cardUpdate := bson.M{
		"$set": bson.M{
			"cardStatus": StatusKeep,
			"updateAt":   time.Now(),
		},
	}
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	collection := client.Database("interviewrbh").Collection("interviewCard")
	updateSuccess, err := collection.UpdateByID(context.Background(), id, cardUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, updateSuccess)
}
