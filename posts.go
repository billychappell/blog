package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

// Post is usedd to unmarshal JSON post data into.
type Post struct {
	Id          int
	Title       string
	CreatedAt   time.Time
	Author      string `sql:"not null"`
	Description string
	Content     string
	Comments    []Comment
	Picture     string
}

// Comment is a reply to a post by a reader.
// IMPORTANT: The Replies field is anonymous,
// which means it is NOT type safe until typecasted.
type Comment struct {
	Id        int
	PostId    int `sql:"index"`
	Content   string
	Author    string `sql:"not null"`
	CreatedAt time.Time
	// []Comment
}

// Posts is a slice of Post structs.
type Posts []Post

var db gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "user=billy password=*********")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Post{}, &Comment{})
}

func samplePost() []Post {
	post := Post{
		Title:       "Test Title",
		CreatedAt:   time.Now(),
		Author:      "Billy Chappell",
		Description: "Test description.",
		Content:     "Test content! Hello world!",
	}
	p := []Post{post}
	return p
}
