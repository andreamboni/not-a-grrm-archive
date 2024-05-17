package router

import (
	"github.com/andreamboni/not-a-grrm-archive/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/not-a-grrm-archive/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/fetchAllBlogPosts", handler.FetchAllBlogPosts)
		v1.POST("/saveBlogPost", handler.SaveBlogPostHandler)
	}
}
