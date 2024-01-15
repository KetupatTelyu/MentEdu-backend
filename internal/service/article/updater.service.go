package article

import (
	"context"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/repository"
	"mentedu-backend/utils"
)

type ArticleUpdaterUsecase interface {
	UpdateArticle(ctx context.Context, id int, title, body, slug, image string, categoryID int, updatedBy string) (*model.Article, error)

	UpdateCategory(ctx context.Context, id int, name, updatedBy string) error
}

type ArticleUpdater struct {
	articleRepo  repository.ArticleRepositoryUseCase
	categoryRepo repository.CategoryRepositoryUseCase
}

func NewArticleUpdater(articleRepo repository.ArticleRepositoryUseCase, categoryRepo repository.CategoryRepositoryUseCase) ArticleUpdaterUsecase {
	return &ArticleUpdater{
		articleRepo:  articleRepo,
		categoryRepo: categoryRepo,
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

func (au *ArticleUpdater) UpdateCategory(ctx context.Context, id int, name, updatedBy string) error {
	category, err := au.categoryRepo.GetById(ctx, id)

	if err != nil {
		return err
	}

	category.Name = name
	category.UpdatedBy = utils.StringToNullString(updatedBy)

	err = au.categoryRepo.Update(ctx, category)

	if err != nil {
		return err
	}

	return nil
}
