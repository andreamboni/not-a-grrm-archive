package handler

import (
	"context"
	"fmt"

	"github.com/andreamboni/not-a-grrm-archive/config"
	"github.com/andreamboni/not-a-grrm-archive/model"
	"github.com/gin-gonic/gin"
)

func SaveBlogPostHandler(ctx *gin.Context) {
	request := CreateBlogPostRequest{}
	ctx.BindJSON(&request)

	blogPost := model.BlogPost{
		PostImage: request.PostImage,
		Title:     request.Title,
		URL:       request.URL,
		TheDate:   request.TheDate,
		Content:   request.Content,
	}

	coll, err := config.GetMongoDBCollection()

	coll.InsertOne(context.TODO(), blogPost)

	if err != nil {
		panic(err)
	}

	fmt.Println("Not a GRRM Archive: " + blogPost.Title + ": " + blogPost.TheDate)
}
