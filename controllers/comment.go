package controllers

import (
	"github.com/musukwamoshi/golang-test-api/databaseconfig"
	"github.com/musukwamoshi/golang-test-api/models"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm"
)

func GetComments(c *fiber.Ctx) error {
	db := databaseconfig.DB
	Comments := []models.Comment{}
	db.Model(&models.Comment{}).Order("ID asc").Limit(100).Find(&Comments)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    Comments,
	})
}
