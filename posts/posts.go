package posts

import (
	"errors"
	"time"
)

type BlogPost struct {
	Title   string      `json:"title"`
	Body    string      `json:"body"`
	Date    time.Time   `json:"created"`
	Id      int         `json:"id"`
}
var posts = make([]BlogPost, 1, 10)

func New(title, body string) BlogPost {
	return BlogPost{
		Title:title,
		Body:body,
		Date: time.Now(),
	}
}

func GetBlogPosts() []BlogPost{
	return posts[:]
}

func CreateNewPost(title, body string) (bool, error) {
	if title == "" {
		return false, errors.New("title must not be empty")
	}
	if body == "" {
		return false, errors.New("body must not be empty")
	}

	idx := 0
	if len(posts) > 0 {
		idx = len(posts)-1
	}
	posts[idx] = New(title, body)
	return true, nil
}
