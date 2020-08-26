package main

import (
	"net/http"
	"strconv"

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

	r.GET("/posts/:id", GetPostByIdHandler())

	r.POST("/posts", NewPostHandler())

	_ = r.Run()
}

func GetPostByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			c.String(http.StatusNotFound, "cannot find post by the specified id")
			return
		}

		p, err := posts.GetById(id)
		if err != nil {
			c.String(http.StatusNotFound, "")
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func NewPostHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		var post Post
		if err := c.BindJSON(&post); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		if _, err := posts.Save(post.Title, post.Body); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusCreated, "created")
	}
}


func GetPostsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		page := c.Query("page")
		p, err := strconv.Atoi(page)
		if err != nil {
			p = 1
		}
		c.JSON(http.StatusOK, posts.All(p))
	}
}


