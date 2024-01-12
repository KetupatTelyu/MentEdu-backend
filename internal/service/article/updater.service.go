package article

import (
	"context"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/repository"
	"mentedu-backend/utils"
)

type ArticleUpdaterUsecase interface {
	UpdateArticle(ctx context.Context, id int, title, body, slug, image string, categoryID int, updatedBy string) (*model.Article, error)
}

type ArticleUpdater struct {
	articleRepo repository.ArticleRepositoryUseCase
}

func NewArticleUpdater(articleRepo repository.ArticleRepositoryUseCase) ArticleUpdaterUsecase {
	return &ArticleUpdater{
		articleRepo: articleRepo,
	}
}

func (au *ArticleUpdater) UpdateArticle(ctx context.Context, id int, title, body, slug, image string, categoryID int, updatedBy string) (*model.Article, error) {
	article, err := au.articleRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	article.Title = title
	article.Body = body
	article.Slug = slug
	article.Image = image
	article.UpdatedBy = utils.StringToNullString(updatedBy)

	article.ArticleCategory = model.NewArticleCategory(article.ID, categoryID, updatedBy)

	err = au.articleRepo.Update(ctx, article)

	if err != nil {
		return nil, err
	}

	return article, nil
}
