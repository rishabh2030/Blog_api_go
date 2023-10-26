package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// model
type Post struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Post{})

	e := echo.New()

	e.POST("/posts", func(c echo.Context) error { return CreatePost(c, db) })
	e.GET("/posts", func(c echo.Context) error { return GetPosts(c, db) })

	e.Logger.Fatal(e.Start(":8080"))

}

func CreatePost(c echo.Context, db *gorm.DB) error {
	post := &Post{}
	c.Bind(post)

	db.Create(post)

	return c.JSON(http.StatusCreated, post)
}

func GetPosts(c echo.Context, db *gorm.DB) error {
	var posts []Post

	db.Find(&posts)

	return c.JSON(http.StatusOK, posts)
}
