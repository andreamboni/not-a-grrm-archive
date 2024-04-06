package handler

import (
	"fmt"

	"github.com/andreamboni/not-a-grrm-archive/model"
	"github.com/gin-gonic/gin"
)

func SaveBlogPostHandler(context *gin.Context) {
	request := CreateBlogPostRequest{}
	context.BindJSON(&request)

	blogPost := model.BlogPost{
		PostImage: request.PostImage,
		Title:     request.Title,
		URL:       request.URL,
		TheDate:   request.TheDate,
		Content:   request.Content,
	}

	fmt.Println(blogPost)
}
