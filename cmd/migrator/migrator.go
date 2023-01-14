package main

import (
	"log"

	"github.com/n0c1337/linkshortener/internal/database"
	"github.com/n0c1337/linkshortener/internal/models"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("[Databse] failed to connect: %s", err.Error())
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Link{})
}
