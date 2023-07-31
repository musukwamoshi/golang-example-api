package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/musukwamoshi/golang-test-api/controllers"
	"github.com/musukwamoshi/golang-test-api/middleware"
)

func Initalize(router *fiber.App) {

	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	users := router.Group("/users")
	users.Post("/", controllers.CreateUser)
	users.Delete("/", middleware.Authenticated, controllers.DeleteUser)
	users.Put("/", middleware.Authenticated, controllers.ChangePassword)
	//users.Post("/profile", middleware.Authenticated, controllers.GetUserInfo)
	users.Post("/login", controllers.Login)
	users.Delete("/logout", controllers.Logout)

	articles := router.Group("/articles", middleware.Authenticated)
	articles.Post("/all", controllers.GetArticles)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
