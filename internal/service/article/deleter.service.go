package article

import (
	"context"
	"mentedu-backend/internal/repository"
)

type ArticleDeleterUsecase interface {
	DeleteArticle(ctx context.Context, id int) error

	DeleteCategory(ctx context.Context, id int) error

	DeleteArticleCategory(ctx context.Context, articleID, categoryID int) error
}

type ArticleDeleter struct {
	articleRepo         repository.ArticleRepositoryUseCase
	categoryRepo        repository.CategoryRepositoryUseCase
	articleCategoryRepo repository.ArticleCategoryRepositoryUseCase
}

func NewArticleDeleter(articleRepo repository.ArticleRepositoryUseCase, categoryRepo repository.CategoryRepositoryUseCase, articleCategoryRepo repository.ArticleCategoryRepositoryUseCase) ArticleDeleterUsecase {
	return &ArticleDeleter{
		articleRepo:         articleRepo,
		categoryRepo:        categoryRepo,
		articleCategoryRepo: articleCategoryRepo,
	}
}
func (ad *ArticleDeleter) DeleteArticle(ctx context.Context, id int) error {
	if err := ad.articleRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ad *ArticleDeleter) DeleteCategory(ctx context.Context, id int) error {
	if err := ad.categoryRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (ad *ArticleDeleter) DeleteArticleCategory(ctx context.Context, articleID, categoryID int) error {
	if err := ad.articleCategoryRepo.Delete(ctx, articleID, categoryID); err != nil {
		return err
	}

	return nil
}
