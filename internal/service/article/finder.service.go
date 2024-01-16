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

	FindCategoryByID(ctx context.Context, id int) (*model.Category, error)
	FindAllCategory(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Category, error)

	FindArticleCategoryByArticleID(ctx context.Context, articleID int) ([]*model.ArticleCategory, error)

	FindArticleByCategoryID(ctx context.Context, categoryID int) ([]*model.Article, error)
}

type ArticleFinder struct {
	articleRepo         repository.ArticleRepositoryUseCase
	categoryRepo        repository.CategoryRepositoryUseCase
	articleCategoryRepo repository.ArticleCategoryRepositoryUseCase
}

func NewArticleFinder(articleRepo repository.ArticleRepositoryUseCase, categoryRepo repository.CategoryRepositoryUseCase, articleCategoryRepo repository.ArticleCategoryRepositoryUseCase) ArticleFinderUsecase {
	return &ArticleFinder{
		articleRepo:         articleRepo,
		categoryRepo:        categoryRepo,
		articleCategoryRepo: articleCategoryRepo,
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

func (af *ArticleFinder) FindCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	category, err := af.categoryRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (af *ArticleFinder) FindAllCategory(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Category, error) {
	categories, err := af.categoryRepo.GetAll(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (af *ArticleFinder) FindArticleCategoryByArticleID(ctx context.Context, articleID int) ([]*model.ArticleCategory, error) {
	articleCategories, err := af.articleCategoryRepo.GetByArticleID(ctx, articleID)

	if err != nil {
		return nil, err
	}

	return articleCategories, nil
}

func (af *ArticleFinder) FindArticleByCategoryID(ctx context.Context, categoryID int) ([]*model.Article, error) {
	articleCategories, err := af.FindArticleCategoryByArticleID(ctx, categoryID)

	if err != nil {
		return nil, err
	}

	var articleIDs []int

	for _, articleCategory := range articleCategories {
		articleIDs = append(articleIDs, articleCategory.ArticleID)
	}

	var articles []*model.Article

	for _, articleID := range articleIDs {
		article, err := af.articleRepo.GetById(ctx, articleID)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}
