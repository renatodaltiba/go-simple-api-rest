package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	URL  string `json:"url"`
}

func CreateBookmark(name, url string) (Bookmark, error) {
	var newBookmark = Bookmark{Name: name, URL: url}

	db, err := gorm.Open(sqlite.Open("bookmarks.db"), &gorm.Config{})
	if err != nil {
		return newBookmark, err
	}
	db.Create(&Bookmark{Name: name, URL: url})

	return newBookmark, nil
}
