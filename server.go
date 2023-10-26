package main

import (
	"blog_api/db"
	"blog_api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	dbInstance := db.InitDB()

	e := echo.New()

	// Initialize handlers
	postHandler := handlers.NewPostHandler(dbInstance)

	e.POST("/posts", postHandler.CreatePost)
	e.GET("/posts", postHandler.GetPosts)

	e.Logger.Fatal(e.Start(":8080"))
}
