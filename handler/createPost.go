package handler

import (
	"fmt"

	"github.com/andreamboni/not-a-grrm-archive/model"
	"github.com/gin-gonic/gin"
)

func CreatePostHandler(context *gin.Context) {
	request := CreatePostRequest{}
	context.BindJSON(&request)

	post := model.Post{
		Image:   request.Image,
		Title:   request.Title,
		URL:     request.URL,
		Date:    request.Date,
		Content: request.Content,
	}

	fmt.Println(post)
}
