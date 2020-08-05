package infra

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/petegabriel/personalblog/posts"
)

var pool *pgxpool.Pool

//init pool
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading .env file: %v\n", err)
		os.Exit(1)
	}

	pp, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URI"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}else {
		pool = pp
	}
}

/**
Fetch connection from pool.
Release method must be called over the resulting connection instance.
 */
func fetch() *pgxpool.Conn {
	c, err := pool.Acquire(context.Background())
	if err != nil{
		fmt.Fprintf(os.Stderr, "Unable to fetch a connection from pool: %v\n", err)
		os.Exit(1)
	}
	return c
}

/**
Save a post in the database.
 */
func Save(pst posts.BlogPost) (int, error){
	//validate fields
	if strings.Compare(pst.Title, "") == 0 {
		log.Println("tried to save a post without field title")
		return -1, errors.New("a post must have a title")
	}
	if strings.Compare(pst.Body, "") == 0 {
		log.Println("tried to save a post without field body")
		return -1, errors.New("a post must have a body")
	}

	con := fetch()
	defer con.Conn().Close(context.Background())


	qry := `INSERT INTO posts (title, body, date) 
                     VALUES ($1, $2, $3) 
                     RETURNING id`
	var id = 0
	err := con.Conn().QueryRow(context.Background(), qry, pst.Title, pst.Body, pst.Date).Scan(&id)
	if err != nil {
		return -1, err
	}

	fmt.Println("New record ID is:", id)
	return id, nil
}

/**
Get a post by its id.
 */
func Get(i int) (*posts.BlogPost, error) {
	qry := `select title, body, id, date from posts where id=$1;`
	var body, title string
	var id int
	var date time.Time

	con := fetch()
	defer con.Conn().Close(context.Background())

	row := con.Conn().QueryRow(context.Background(), qry, i)

	if err:= row.Scan(&title, &body, &id, &date); err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("post with id %d not found", i))
	}
	return &posts.BlogPost{
		Title: title,
		Body:  body,
		Id:    id,
		Date:  date,
	}, nil
}