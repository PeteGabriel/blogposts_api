package posts

import (
	"testing"

	"github.com/matryer/is"
)

func TestCreateNewPost(t *testing.T){
	i := is.New(t)
	title := "name"
	body := "message"
	created, err := Save(title, body)
	i.NoErr(err)
	i.True(created)
}

func TestCreateNewPost_WithEmptyTitle(t *testing.T){
	i := is.New(t)
	created, err := Save("", "body")
	i.True(!created)
	i.True(err != nil)
}

func TestCreateNewPost_WithEmptyBody(t *testing.T){
	i := is.New(t)
	created, err := Save("title", "")
	i.True(!created)
	i.True(err != nil)
}

func TestCreateNewPost_WithEmptyDatay(t *testing.T){
	i := is.New(t)
	created, err := Save("", "")
	i.True(!created)
	i.True(err != nil)
}