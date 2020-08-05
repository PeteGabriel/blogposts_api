package infra

import (
	"context"
	"testing"

	is2 "github.com/matryer/is"
	"github.com/petegabriel/personalblog/posts"
)


func TestSavePostWithInvalidData(t *testing.T){
	is := is2.New(t)
	post := posts.New("Test Post", "")

	id, err := Save(post)
	is.True(err != nil)
	is.True(err.Error() == "a post must have a body")
	is.True(id < 0)

	p, err := Get(id)
	is.True(err != nil)
	is.True(err.Error() == "post with id -1 not found")
	is.True(p == nil)
}

func TestSavePost(t *testing.T) {
	is := is2.New(t)

	post := posts.New("Test Post", "This post is a test")
	id, err := Save(post)
	is.NoErr(err)
	is.True(id > 0)

	p, err := Get(id)
	is.NoErr(err)
	is.True(id == p.Id)
	is.True(post.Title == p.Title)
	is.True(post.Body == p.Body)
}

func TestGetWithNoElements(t *testing.T){
	is := is2.New(t)
	con := fetch()
	qry := "DELETE FROM posts"
	_ = con.Conn().QueryRow(context.Background(), qry).Scan()

	p, err := Get(1)
	is.True(err != nil)
	is.True(err.Error() == "post with id 1 not found")
	is.True(p == nil)
}
