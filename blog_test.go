package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	is2 "github.com/matryer/is"
)

var router *gin.Engine

func init(){
	router = gin.Default()
	router.GET("/posts", GetPostsHandler())
	router.POST("/posts", NewPostHandler())
}

func TestGetPostsHandler(t *testing.T){
	is := is2.New(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts", nil)
	router.ServeHTTP(w, req)

	is.Equal(200, w.Code)
	is.Equal("[{\"title\":\"\",\"body\":\"\",\"created\":\"0001-01-01T00:00:00Z\",\"id\":0}]", w.Body.String())
}

func TestNewPostHandler(t *testing.T){
	is := is2.New(t)
	w := httptest.NewRecorder()
	body := &Post{Title:"My First Post", Body:"It goes like this..."}
	reader, err := json.Marshal(body)
	is.NoErr(err)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewReader(reader))
	router.ServeHTTP(w, req)
	is.Equal(201, w.Code)
	is.Equal("created", w.Body.String())

}

func TestNewPostHandler_WithIncorrectBody(t *testing.T){
	is := is2.New(t)
	w := httptest.NewRecorder()
	body := &Post{Title:"My First Post"}
	reader, err := json.Marshal(body)
	is.NoErr(err)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewReader(reader))
	router.ServeHTTP(w, req)
	is.Equal(400, w.Code)
}

