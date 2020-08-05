package infra

import (
	"context"
	"strconv"
	"testing"

	is2 "github.com/matryer/is"
)


func TestSavePostWithInvalidData(t *testing.T){
	is := is2.New(t)

	id, err := Save("Test Post", "")
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

	post := New("Test Post", "This post is a test")
	id, err := Save("Test Post", "This post is a test")
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
	defer con.Conn().Close(context.Background())
	_ = con.Conn().QueryRow(context.Background(), qry).Scan()

	p, err := Get(1)
	is.True(err != nil)
	is.True(err.Error() == "post with id 1 not found")
	is.True(p == nil)
}

func TestGetAllElements(t *testing.T) {
	is := is2.New(t)
	con := fetch()
	defer con.Conn().Close(context.Background())
	qry := "DELETE FROM posts"
	_ = con.Conn().QueryRow(context.Background(), qry).Scan()

	for i := 0; i < 10; i++ {
		tb := strconv.Itoa(i+1)
		_, err := Save(tb, tb)
		if err != nil {

		}
	}

	posts := All()
	is.True(len(posts) == 10)
}