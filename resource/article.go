package resource

import "mime/multipart"

type CreateArticleRequest struct {
	Title      string                `json:"title" binding:"required"`
	Body       string                `json:"body" binding:"required"`
	Slug       string                `json:"slug" binding:"required"`
	Image      *multipart.FileHeader `json:"image" binding:"required"`
	CategoryID int                   `json:"category_id" binding:"required"`
}

type UpdateArticleRequest struct {
	ID         string                `json:"id" binding:"required"`
	Title      string                `json:"title" binding:"required"`
	Body       string                `json:"body" binding:"required"`
	Slug       string                `json:"slug" binding:"required"`
	Image      *multipart.FileHeader `json:"image" binding:"required"`
	CategoryID int                   `json:"category_id" binding:"required"`
}
