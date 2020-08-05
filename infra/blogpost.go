package infra

import (
	"time"
)

type BlogPost struct {
	Title   string      `json:"title"`
	Body    string      `json:"body"`
	Id      int         `json:"id"`
	Date    time.Time   `json:"date"`
}

/**
New post instance
*/
func New(title, body string) BlogPost {
	return BlogPost{
		Title:title,
		Body:body,
		Date: time.Now().UTC(),
	}
}