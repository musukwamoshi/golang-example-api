package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/musukwamoshi/golang-test-api/controllers"
	model "github.com/musukwamoshi/golang-test-api/models"
)

func Authenticated(c *fiber.Ctx) error {
	json := new(model.Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Session Format",
		})
	}
	user, err := controllers.GetUser(json.ID)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "404: not found",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
