package responses

import (
	"mentedu-backend/internal/model"
)

type ArticleResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	Image       string `json:"image"`
	Category    string `json:"category"`
	CreatedBy   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
}

func FromArticle(article *model.Article) *ArticleResponse {
	return &ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Slug:        article.Slug,
		Body:        article.Body,
		Image:       article.Image,
		CreatedBy:   article.CreatedBy.String,
		CreatedDate: article.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
