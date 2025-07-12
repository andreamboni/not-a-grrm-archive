package handler

import (
	"context"
	"net/http"

	"github.com/andreamboni/not-a-grrm-archive/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAllBlogPosts(ctx *gin.Context) {

	logger := config.NewLogger("fetch all blogposts")
	coll, err := config.GetMongoDBCollection()

	if err != nil {
		logger.Errorf("error trying to get db configs: %v", err)
	}

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		logger.Errorf("error trying to find posts: %v", err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		SendError(ctx, http.StatusNotFound, "no blog posts were found")
	}

	if len(results) > 0 {
		SendSuccess(ctx, "blog posts retrieved successfully", results)
	}

}
