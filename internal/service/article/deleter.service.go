package article

import (
	"context"
	"mentedu-backend/internal/repository"
)

type ArticleDeleterUsecase interface {
	DeleteArticle(ctx context.Context, id int) error
}

type ArticleDeleter struct {
	articleRepo repository.ArticleRepositoryUseCase
}

func NewArticleDeleter(articleRepo repository.ArticleRepositoryUseCase) ArticleDeleterUsecase {
	return &ArticleDeleter{
		articleRepo: articleRepo,
	}
}

func (ad *ArticleDeleter) DeleteArticle(ctx context.Context, id int) error {
	if err := ad.articleRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
