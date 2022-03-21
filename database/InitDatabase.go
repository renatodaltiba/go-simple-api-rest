package database

import (
	"github.com/renatodaltiba/go-rest-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("bookmarks.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&model.Bookmark{})

	return nil
}
