package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

// Post is used to store fields and values from the "posts" table
// in a Golang struct.
type Post struct {
	ID          int
	Title       string
	CreatedAt   time.Time
	Author      string `sql:"not null"`
	Description string
	Content     string
	ImageURL    string
	Path        string
	Featured    bool
}

// Posts is an array of pointers to a post, which is used to create
// handlers and generate pages for each post retrieved from the database.
type Posts []*Post

// Featured is used to show a post at the top of the index indefinitely.
type Featured *Post

// Data holds the Posts array and featured article so we can easily pass
// them to templates.
type Data struct {
	P Posts
	F Featured
}

func open() *sql.DB {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=testing host=172.17.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Gets a single post by its id
func retrievePost(id int, db *sql.DB) (p *Post, err error) {
	err = db.QueryRow("SELECT id, title, createdat, content, author, description, imageurl, path, featured FROM posts where id = $1",
		id).Scan(&p.Title, &p.CreatedAt, &p.Content, &p.Author, &p.Description, &p.ImageURL, &p.Path, &p.Featured)
	return
}

// Create a new post
func (p *Post) Create(db *sql.DB) (err error) {
	statement := `insert into posts (title, createdat, content, author, description, imageurl, path, featured)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Title, p.CreatedAt, p.Content, p.Author, p.Description, p.ImageURL, p.Path, p.Featured).Scan(&p.ID)
	return err
}

// Get a list of all posts and their properties for the index.
func GetPosts() (ps Posts, err error) {
	db := open()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, title, createdat, content, author, description, imageurl, path, featured FROM posts ORDER BY createdat DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		p := Post{}
		if err = rows.Scan(&p.ID, &p.Title, &p.CreatedAt, &p.Content, &p.Author, &p.Description, &p.ImageURL, &p.Path, &p.Featured); err != nil {

			return
		}
		ps = append(ps, &p)
	}
	rows.Close()
	return ps, err
}
