package posts

import (
	"errors"
	"time"
)

type BlogPost struct {
	Title   string      `json:"title"`
	Body    string      `json:"body"`
	Id      int       `json:"id"`
	Date    time.Time   `json:"date"`
}
var posts = make([]BlogPost, 0)

func New(title, body string) BlogPost {
	return BlogPost{
		Title:title,
		Body:body,
		Date: time.Now().UTC(),
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

	if posts == nil {
		posts = make([]BlogPost, 0)
	}
	posts = append(posts, New(title, body))
	return true, nil
}
