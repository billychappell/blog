package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	DB_USER     = flag.String("db-user", "billy", "database username to login with")
	DB_PASSWORD = flag.String("db-pass", "", "database password to login with")
	db          *sql.DB
)

type Post struct {
	ID          int
	Title       string
	CreatedAt   time.Time
	Author      string `sql:"not null"`
	Description string
	Content     string
	Comments    []Comment
	ImageURL    string
}

// Comment is a reply to a post by a reader.
// IMPORTANT: The Replies field is anonymous,
// which means it is NOT type safe until typecasted.
type Comment struct {
	ID        int
	Post      *Post
	Content   string
	Author    string `sql:"not null"`
	CreatedAt time.Time
	// []Comment
}

// Posts is a slice of Post structs.
type Posts []Post

func init() {
	dbInfo := fmt.Sprintf("user=%s password=%s sslmode=disable", DB_USER, DB_PASSWORD)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Starting database, info: \n %v\n", db)
}

func (p *Post) Create() error {
	err := db.QueryRow(`insert into posts (title, author, description, content, createdat, imageurl)
			values ($1, $2, $3, $4, $5, $6) returning id`, p.Title, p.Author, p.Description, p.Content,
		p.CreatedAt, p.ImageURL).Scan(&p.ID)
	return err
}

func GetPost(id int) (p Post, err error) {
	p = Post{}
	p.Comments = []Comment{}
	err = db.QueryRow("SELECT id, title, author, description, content, createdat, imageurl FROM posts WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Author, &p.Description, &p.Content, &p.CreatedAt, &p.ImageURL)

	rows, err := db.Query("SELECT id, content, author, createdat FROM comments WHERE post_id = $1, id")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &p}
		err = rows.Scan(&comment.ID, &comment.Content, &comment.Author, &comment.CreatedAt)
		if err != nil {
			return
		}
		p.Comments = append(p.Comments, comment)
	}
	rows.Close()
	return
}

func samplePost() []Post {
	post := Post{
		Title:       "Test Title",
		CreatedAt:   time.Now(),
		Author:      "Billy Chappell",
		Description: "Test description.",
		Content:     "Test content! Hello world!",
		ImageURL:    "http://placehold.it/1024x768",
	}
	p := []Post{post}
	return p
}
