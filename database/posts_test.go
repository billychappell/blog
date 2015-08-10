package database

import (
	"fmt"
	"testing"
	"time"
)

func TestCreatePost(t *testing.T) {
	db := opendb()
	post := Post{
		Title:       "Test posting",
		CreatedAt:   time.Now(),
		Author:      "Billy Chappell",
		Description: "A test post.",
		Content:     "Test test test test test test test",
	}

	if err := post.Create(db); err != nil {
		t.Errorf("couldn't create post %s: \n %v\n", post.Title, err)
	}
}

func TestRetrievePost(t *testing.T) {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	db := opendb()
	rows, err := db.Query("SELECT id FROM posts WHERE title = $1", "Test posting")
	if err != nil {
		t.Errorf("couldn't retrieve test post with title %s: \n %v \n", "Test posting", err)
	}
	fmt.Println(rows)
}
