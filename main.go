package main

import (
	"interview-rbh/model"
	"interview-rbh/service/interview"
	"interview-rbh/service/login"

	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB()
	login.InitRateLimit()
	login.CreateUser()
	interview.CreateInterview()
	r := gin.Default()
	r.POST("/login", login.Login)
	cardGroup := r.Group("card")
	cardGroup.Use(login.Authenticate())
	cardGroup.Use(login.LimitMiddleware())
	route(cardGroup)

	r.Run()
}

func route(r *gin.RouterGroup) {
	r.GET("/fetch", interview.FetchCard)
	r.GET("/view/:cardId", interview.ViewCard)
	r.POST("/comment/add", interview.AddComment)
	r.POST("/comment/update", interview.EditComment)
	r.POST("/comment/remove", interview.DeleteComment)
	r.POST("/update", interview.UpdateCard)
	r.GET("/history/:cardId", interview.GetHistoryByCardID)
	r.GET("/keep/:cardId", interview.KeepInterviewCard)
}
