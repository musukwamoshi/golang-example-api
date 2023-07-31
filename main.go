package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/musukwamoshi/golang-test-api/databaseconfig"
	"github.com/musukwamoshi/golang-test-api/router"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, nikschaefer.tech"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	databaseconfig.ConnectDB()

	router.Initalize(app)

	// app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": "Welcome to Golang, Fiber, and GORM",
	// 	})
	// })

	log.Fatal(app.Listen(":8000"))
}
