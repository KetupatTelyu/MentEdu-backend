package model

type Article struct {
	ID              int              `json:"id" form:"title" gorm:"primaryKey;autoIncrement"`
	Title           string           `json:"title" form:"title"`
	Body            string           `json:"body" form:"body" gorm:"type:text"`
	Image           string           `json:"image" form:"image" gorm:"type:text"`
	Slug            string           `json:"slug" form:"slug"`
	ArticleCategory *ArticleCategory `json:"article_category"`
	Auditable
}

func NewArticle(
	title string,
	body string,
	slug string,
	image string,
	createdBy string,
) *Article {
	return &Article{
		Title:     title,
		Body:      body,
		Slug:      slug,
		Image:     image,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *Article) MapUpdateFrom(from *Article) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"title":      model.Title,
			"body":       model.Body,
			"slug":       model.Slug,
			"image":      model.Image,
			"created_by": model.CreatedBy,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Title != from.Title {
		mapped["title"] = from.Title
	}

	if model.Body != from.Body {
		mapped["body"] = from.Body
	}

	if model.Slug != from.Slug {
		mapped["slug"] = from.Slug
	}

	if model.Image != from.Image {
		mapped["image"] = from.Image
	}

	mapped["updated_at"] = model.UpdatedAt
	return &mapped
}
