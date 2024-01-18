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
	CreateCategory(ctx context.Context, name, createdBy string) error
	CreateArticleCategory(ctx context.Context, articleID, categoryID int, createdBy string) error
	CreateConsultant(ctx context.Context, consultant *model.Consultant) error
}

type ArticleCreator struct {
	cfg                 config.Config
	articleRepo         repository.ArticleRepositoryUseCase
	categoryRepo        repository.CategoryRepositoryUseCase
	articleCategoryRepo repository.ArticleCategoryRepositoryUseCase
	cloudStorage        utils.CloudStorage
}

func NewArticleCreator(cfg config.Config, articleRepo repository.ArticleRepositoryUseCase, categoryRepo repository.CategoryRepositoryUseCase, articleCategoryRepo repository.ArticleCategoryRepositoryUseCase, cloudStorage utils.CloudStorage) ArticleCreatorUsecase {
	return &ArticleCreator{
		cfg:                 cfg,
		articleRepo:         articleRepo,
		categoryRepo:        categoryRepo,
		articleCategoryRepo: articleCategoryRepo,
		cloudStorage:        cloudStorage,
	}
}

func (ac *ArticleCreator) CreateArticle(ctx context.Context, title, body, slug, image string, categoryID int, createdBy string) (*model.Article, error) {
	article := model.NewArticle(title, body, slug, image, createdBy)
	article, err := ac.articleRepo.Create(ctx, article)
	if err != nil {
		return nil, err
	}

	articleCategory := model.NewArticleCategory(article.ID, categoryID, createdBy)

	err = ac.articleCategoryRepo.Create(ctx, articleCategory)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (ac *ArticleCreator) CreateCategory(ctx context.Context, name, createdBy string) error {
	category := model.NewCategory(0, name, createdBy)
	err := ac.categoryRepo.Create(ctx, category)

	if err != nil {
		return err
	}

	return nil
}

func (ac *ArticleCreator) CreateArticleCategory(ctx context.Context, articleID, categoryID int, createdBy string) error {
	articleCategory := model.NewArticleCategory(articleID, categoryID, createdBy)
	err := ac.articleCategoryRepo.Create(ctx, articleCategory)

	if err != nil {
		return err
	}

	return nil
}

func (ac *ArticleCreator) CreateConsultant(ctx context.Context, consultant *model.Consultant) error {
	return nil
}
