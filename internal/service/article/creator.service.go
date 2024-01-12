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

func NewArticleCreator(cfg config.Config, articleRepo repository.ArticleRepositoryUseCase, cloudStorage utils.CloudStorage) ArticleCreatorUsecase {
	return &ArticleCreator{
		cfg:          cfg,
		articleRepo:  articleRepo,
		cloudStorage: cloudStorage,
	}
}

func (ac *ArticleCreator) CreateArticle(ctx context.Context, title, body, slug, image string, categoryID int, createdBy string) (*model.Article, error) {
	article := model.NewArticle(title, body, slug, image, createdBy)

	article, err := ac.articleRepo.Create(ctx, article)

	if err != nil {
		return nil, err
	}

	article.ArticleCategory = model.NewArticleCategory(article.ID, categoryID, createdBy)

	err = ac.articleRepo.Update(ctx, article)

	if err != nil {
		return nil, err
	}

	return article, nil
}
