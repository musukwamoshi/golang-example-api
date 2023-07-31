package controllers

import (
	"github.com/musukwamoshi/golang-test-api/databaseconfig"
	"github.com/musukwamoshi/golang-test-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	db := databaseconfig.DB
	Articles := []models.Article{}
	db.Model(&models.Article{}).Order("ID asc").Limit(100).Find(&Articles)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    Articles,
	})
}
