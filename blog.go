package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petegabriel/personalblog/posts"
)

type Post struct {
	Title string `form:"title" json:"title" binding:"required"`
	Body  string `form:"body" json:"body" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/posts", GetPostsHandler())

	r.POST("/posts", NewPostHandler())

	_ = r.Run()
}

func NewPostHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var post Post
		if err := c.BindJSON(&post); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		created, err := posts.CreateNewPost(post.Title, post.Body)

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		if created {
			c.String(http.StatusCreated, "created")
		}
	}
}

func GetPostsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, posts.GetBlogPosts())
	}
}


