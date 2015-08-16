package database

import (
	"fmt"
	"testing"
	"time"
)

func TestCreatePost(t *testing.T) {
	db := open()
	post := Post{
		Title:       "Test posting",
		CreatedAt:   time.Now(),
		Author:      "Billy Chappell",
		Description: "A test post.",
		Content:     "Test test test test test test test",
		ImageURL:    "static/test.png",
		Path:        "test",
		Featured:    false,
	}

	if err := post.Create(db); err != nil {
		t.Errorf("couldn't create post %s: \n %v\n", post.Title, err)
	}
}

func TestRetrievePost(t *testing.T) {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	db := open()
	rows, err := db.Query("SELECT id FROM posts WHERE title = $1", "Test posting")
	if err != nil {
		t.Errorf("couldn't retrieve test post with title %s: \n %v \n", "Test posting", err)
	}
	fmt.Println(rows)
}
