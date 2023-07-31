package controllers

import (
	"github.com/musukwamoshi/golang-test-api/databaseconfig"
	"github.com/musukwamoshi/golang-test-api/models"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm"
)

func GetReplies(c *fiber.Ctx) error {
	db := databaseconfig.DB
	Replies := []models.Reply{}
	db.Model(&models.Comment{}).Order("ID asc").Limit(100).Find(&Replies)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    Replies,
	})
}
