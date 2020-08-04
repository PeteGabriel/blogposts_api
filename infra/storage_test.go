package infra

import (
	"testing"
	"time"

	is2 "github.com/matryer/is"
	"github.com/petegabriel/personalblog/posts"
)

func TestSavePost(t *testing.T) {
	is := is2.New(t)
	post := &posts.BlogPost{
		Title: "Test Post",
		Body:  "This post is a test",
		Date: time.Now().UTC(),
	}
	id, err := Save(post)
	is.NoErr(err)
	is.True(id > 0)

	p, err := Get(id)
	is.NoErr(err)
	is.True(id == p.Id)
	is.True(post.Title == p.Title)
	is.True(post.Body == p.Body)

	//TODO reset state
}
