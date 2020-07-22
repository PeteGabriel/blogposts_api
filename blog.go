package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petegabriel/personalblog/posts"
)

type post struct {
	Title string `json:"title"`
	Body string `json:"body"`
}


func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK,  posts.GetBlogPosts())
	})

	r.POST("/posts", func(c *gin.Context) {
		var post post
		err := c.BindJSON(&post)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		created, err := posts.CreateNewPost(post.Title, post.Body)

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		if created {
			c.String(http.StatusCreated, "created")
		}
	})

	r.Run()
}
