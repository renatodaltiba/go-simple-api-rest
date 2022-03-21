package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renatodaltiba/go-rest-api/model"
)

func GetAllBookmarks(c *fiber.Ctx) error {
	return c.SendString("All Bookmarks")
}

func SaveBookMark(c *fiber.Ctx) error {
	newBookmark := new(model.Bookmark)

	err := c.BodyParser(newBookmark)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := model.CreateBookmark(newBookmark.Name, newBookmark.URL)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Bookmark created",
		"data":    result,
	})

	return nil
}
