package controllers

import (
	"github.com/musukwamoshi/golang-test-api/databaseconfig"
	"github.com/musukwamoshi/golang-test-api/models"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm"
)

// func CreateProduct(c *fiber.Ctx) error {
// 	db := database.DB
// 	json := new(Product)
// 	if err := c.BodyParser(json); err != nil {
// 		return c.JSON(fiber.Map{
// 			"code":    400,
// 			"message": "Invalid JSON",
// 		})
// 	}
// 	user := c.Locals("user").(User)
// 	newProduct := Product{
// 		UserRefer: user.ID,
// 		Name:      json.Name,
// 		Value:     json.Value,
// 	}
// 	err := db.Create(&newProduct).Error
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}
// 	return c.JSON(fiber.Map{
// 		"code":    200,
// 		"message": "sucess",
// 	})
// }
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
