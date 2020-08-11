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
)

var pool *pgxpool.Pool

//init pool
func init() {
	err := godotenv.Load(os.Getenv("ENV_PATH"))
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
func Save(title, body string) (int, error){
	//validate fields
	if strings.Compare(title, "") == 0 {
		log.Println("tried to save a post without field title")
		return -1, errors.New("a post must have a title")
	}
	if strings.Compare(body, "") == 0 {
		log.Println("tried to save a post without field body")
		return -1, errors.New("a post must have a body")
	}

	con := fetch()
	defer con.Release()


	qry := `INSERT INTO posts (title, body, date) 
                     VALUES ($1, $2, $3) 
                     RETURNING id`
	var id = 0
	pst := New(title, body)
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
func Get(i int) (*BlogPost, error) {
	qry := `select title, body, id, date from posts where id=$1;`
	var body, title string
	var id int
	var date time.Time

	con := fetch()
	defer con.Release()

	row := con.Conn().QueryRow(context.Background(), qry, i)

	if err:= row.Scan(&title, &body, &id, &date); err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("post with id %d not found", i))
	}
	return &BlogPost{
		Title: title,
		Body:  body,
		Id:    id,
		Date:  date,
	}, nil
}

func All() []BlogPost {
	qry := "select title, body, id, date from posts;"
	con := fetch()
	defer con.Release()

	rows, err := con.Conn().Query(context.Background(), qry)
	if err != nil {

	}
	var posts []BlogPost
	defer rows.Close()
	for rows.Next() {
		var body, title string
		var id int
		var date time.Time
		if err:= rows.Scan(&title, &body, &id, &date); err != nil {
			log.Println(err)
		}
		p := BlogPost{
			Title: title,
			Body:  body,
			Id:    id,
			Date:  date,
		}
		posts = append(posts, p)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
	}

	return posts
}