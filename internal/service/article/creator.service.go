package article

import (
	"context"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/repository"
	"mentedu-backend/utils"
)

type ArticleCreatorUsecase interface {
	CreateArticle(ctx context.Context, title, body, slug, image string, categoryID int, createdBy string) (*model.Article, error)
}

type ArticleCreator struct {
	cfg          config.Config
	articleRepo  repository.ArticleRepositoryUseCase
	cloudStorage utils.CloudStorage
}

func NewArticleCreator(articleRepo repository.ArticleRepositoryUseCase, cloudStorage utils.CloudStorage) ArticleCreatorUsecase {
	return &ArticleCreator{
		articleRepo:  articleRepo,
		cloudStorage: cloudStorage,
	}
}

func (ac *ArticleCreator) CreateArticle(ctx context.Context, title, body, slug, image string, categoryID int, createdBy string) (*model.Article, error) {
	article := model.NewArticle(title, body, slug, image, createdBy)

	article.ArticleCategory = &model.ArticleCategory{
		ArticleID:  article.ID,
		CategoryID: categoryID,
	}

	if err := ac.articleRepo.Create(ctx, article); err != nil {
		return nil, err
	}

	return article, nil
}
