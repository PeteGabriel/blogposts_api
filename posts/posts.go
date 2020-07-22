package posts

import (
	"errors"
	"time"
)

type BlogPost struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Date time.Time `json:"created"`
}
var posts = make([]BlogPost, 10)

func New(title, body string) BlogPost {
	return BlogPost{
		Title:title,
		Body:body,
		Date: time.Now(),
	}
}

func GetBlogPosts() []BlogPost{
	return posts[:10]
}

func CreateNewPost(title, body string) (bool, error) {
	if title == "" {
		return false, errors.New("title must not be empty")
	}
	if body == "" {
		return false, errors.New("body must not be empty")
	}

	posts[len(posts)-1] = New(title, body)
	return true, nil
}
