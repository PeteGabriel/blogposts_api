package posts

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetBlogPosts(t *testing.T){
	i := is.New(t)
	posts := GetBlogPosts()
	i.Equal(len(posts), 10)
}

func TestCreateNewPost(t *testing.T){
	i := is.New(t)
	title := "name"
	body := "message"
	created, err := CreateNewPost(title, body)
	i.NoErr(err)
	i.True(created)
}

func TestCreateNewPost_WithEmptyTitle(t *testing.T){
	i := is.New(t)
	created, err := CreateNewPost("", "body")
	i.True(!created)
	i.True(err != nil)
}

func TestCreateNewPost_WithEmptyBody(t *testing.T){
	i := is.New(t)
	created, err := CreateNewPost("title", "")
	i.True(!created)
	i.True(err != nil)
}

func TestCreateNewPost_WithEmptyDatay(t *testing.T){
	i := is.New(t)
	created, err := CreateNewPost("", "")
	i.True(!created)
	i.True(err != nil)
}