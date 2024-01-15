package resource

import "mime/multipart"

type CreateArticleRequest struct {
	Title      string                `json:"title" binding:"required" form:"title"`
	Body       string                `json:"body" binding:"required" form:"body"`
	Slug       string                `json:"slug" binding:"required" form:"slug"`
	Image      *multipart.FileHeader `json:"image" binding:"required" form:"image"`
	CategoryID int                   `json:"category_id" binding:"required" form:"category_id"`
}

type UpdateArticleRequest struct {
	ID         string                `json:"id" binding:"required"`
	Title      string                `json:"title" binding:"required"`
	Body       string                `json:"body" binding:"required"`
	Slug       string                `json:"slug" binding:"required"`
	Image      *multipart.FileHeader `json:"image" binding:"required"`
	CategoryID int                   `json:"category_id" binding:"required"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
