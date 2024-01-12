package article

import (
	"context"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/repository"
)

type ArticleFinderUsecase interface {
	FindByID(ctx context.Context, id int) (*model.Article, error)
	FindAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Article, error)
	FindBySlug(ctx context.Context, slug string) (*model.Article, error)
}

type ArticleFinder struct {
	articleRepo repository.ArticleRepositoryUseCase
}

func NewArticleFinder(articleRepo repository.ArticleRepositoryUseCase) ArticleFinderUsecase {
	return &ArticleFinder{
		articleRepo: articleRepo,
	}
}

func (af *ArticleFinder) FindByID(ctx context.Context, id int) (*model.Article, error) {
	article, err := af.articleRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (af *ArticleFinder) FindAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Article, error) {
	articles, err := af.articleRepo.GetAll(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (af *ArticleFinder) FindBySlug(ctx context.Context, slug string) (*model.Article, error) {
	article, err := af.articleRepo.GetBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}

	return article, nil
}
