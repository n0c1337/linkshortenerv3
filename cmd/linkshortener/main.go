package main

import (
	"log"

	"github.com/n0c1337/linkshortener/internal/auth"
	"github.com/n0c1337/linkshortener/internal/config"
	"github.com/n0c1337/linkshortener/internal/database"
	"github.com/n0c1337/linkshortener/internal/webserver"
)

const configFile = "./config.json"

func main() {
	cfg, err := config.ParseConfigFromJSONFile(configFile)
	if err != nil {
		log.Fatalf("[Config] failed to parse config: %s", err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("[Databse] failed to connect: %s", err.Error())
	}
	auth := auth.NewAuthorization()
	ws := webserver.NewWebServer(db, auth, cfg)
	err = ws.ListenAndServe()
	if err != nil {
		log.Fatalf("[WebServer] failed to start: %s", err.Error())
	}
}
