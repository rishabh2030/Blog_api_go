package handlers

import (
	"blog_api/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostHandler struct {
	DB *gorm.DB
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{DB: db}
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	post := &models.Post{}
	fmt.Println("post")
	c.Bind(post)

	h.DB.Create(post)

	return c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetPosts(c echo.Context) error {
	var posts []models.Post

	h.DB.Find(&posts)

	return c.JSON(http.StatusOK, posts)
}
