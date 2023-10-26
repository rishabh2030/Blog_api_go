package main

import (
	"blog_api/db"
	"blog_api/handlers"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Configure logrus
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	dbInstance := db.InitDB()

	e := echo.New()

	// Initialize handlers
	postHandler := handlers.NewPostHandler(dbInstance)

	e.POST("/posts", postHandler.CreatePost)
	e.GET("/posts", postHandler.GetPosts)

	log.Info("Starting the server...") // This will be logged to the console

	e.Logger.Fatal(e.Start(":8080"))
}
