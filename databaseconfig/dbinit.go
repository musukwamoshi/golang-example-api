package databaseconfig

import (
	"log"
	"os"

	model "github.com/musukwamoshi/golang-test-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&model.Article{}, &model.User{}, &model.Comment{}, &model.Reply{}, &model.Session{})
	if err != nil {
		log.Fatal(err)
	}
}
