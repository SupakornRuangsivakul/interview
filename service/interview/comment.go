package interview

import (
	"context"
	"interview-rbh/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentRequest struct {
	CardID    string `json:"cardId"`
	Comment   string `json:"comment"`
	CommentID string `json:"commentId"`
}

type comment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CardID        string             `bson:"cardId" json:"cardId"`
	CommentDetail string             `bson:"commentDetail" json:"commentDetail"`
	CreatedBy     string             `bson:"createBy" json:"createBy"`
	CreatedAt     time.Time          `bson:"createAt" json:"createAt"`
	UpdatedAt     time.Time          `bson:"updateAt" json:"updateAt"`
}

func AddComment(c *gin.Context) {
	user := c.GetString("userId")
	var req commentRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)
	collection := client.Database("interviewrbh").Collection("commentCard")
	now := time.Now()

	comment := comment{
		CardID:        req.CardID,
		CommentDetail: req.Comment,
		CreatedBy:     user,
		CreatedAt:     now,
	}
	insertResult, err := collection.InsertOne(context.Background(), comment)
	if err != nil {
		log.Fatal(gin.H{"err": err.Error()})
	}
	c.JSON(200, insertResult)
}

func EditComment(c *gin.Context) {
	user := c.GetString("userId")
	var req commentRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)
	collection := client.Database("interviewrbh").Collection("commentCard")
	id, err := primitive.ObjectIDFromHex(req.CommentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	filter := bson.M{"_id": id}
	var comment comment
	// Finding the document
	err = collection.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if comment.CreatedBy != user {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can edit with own user"})
		return
	}
	commentUpdate := bson.M{
		"$set": bson.M{
			"commentDetail": req.Comment,
			"updateAt":      time.Now(),
		},
	}
	_, err = collection.UpdateByID(context.Background(), id, commentUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

}

func DeleteComment(c *gin.Context) {
	user := c.GetString("userId")
	var req commentRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)
	collection := client.Database("interviewrbh").Collection("commentCard")
	id, err := primitive.ObjectIDFromHex(req.CommentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	filter := bson.M{"_id": id}
	var comment comment
	// Finding the document
	err = collection.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if comment.CreatedBy != user {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can remove with own user"})
		return
	}
	comment.CommentDetail = req.Comment
	comment.UpdatedAt = time.Now()
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

}
