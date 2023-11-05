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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cardDetailResponse struct {
	ID           string            `json:"id"`
	CardStatus   string            `json:"cardStatus"`
	CardDetail   string            `json:"cardDetail"`
	CardName     string            `json:"cardName"`
	Email        string            `json:"email"`
	CreatedBy    string            `json:"createBy"`
	CreatedAt    string            `json:"createAt"`
	CommentLists []commentResponse `json:"commentLists"`
}

type commentResponse struct {
	CommentID     string
	User          string
	CommentDetail string
	CreatedAt     string
}

func ViewCard(c *gin.Context) {
	idStr := c.Param("cardId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := model.GetMongoClient(ctx)

	result, err := getCardByID(idStr, client, ctx)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	cardResp := cardDetailResponse{
		ID:         result.ID.Hex(),
		CardStatus: result.CardStatus,
		CardDetail: result.CardDetail,
		CardName:   result.CardName,
		Email:      result.Email,
		CreatedBy:  result.CreatedBy,
		CreatedAt:  utils.GetThaiDate(result.CreatedAt),
	}

	//get all comment
	collection := client.Database("interviewrbh").Collection("commentCard")
	filter := bson.M{"cardId": result.ID.Hex()}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "createAt", Value: -1}})
	cur, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(context.Background())
	var comments []comment
	for cur.Next(context.Background()) {
		var elem comment
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, elem)
	}
	var commentRespList []commentResponse
	for _, item := range comments {
		commentResp := commentResponse{
			CommentID:     item.ID.Hex(),
			User:          item.CreatedBy,
			CommentDetail: item.CommentDetail,
			CreatedAt:     utils.GetThaiDate(item.CreatedAt),
		}
		commentRespList = append(commentRespList, commentResp)
	}
	cardResp.CommentLists = commentRespList
	c.JSON(200, cardResp)
}

func getCardByID(idStr string, client *mongo.Client, ctx context.Context) (interviewCard, error) {
	var result interviewCard

	collection := client.Database("interviewrbh").Collection("interviewCard")
	// Convert the string ID to a MongoDB ObjectID
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": id}

	// Finding the document
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
