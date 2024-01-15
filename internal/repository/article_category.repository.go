package repository

import (
	"context"
	"gorm.io/gorm"
	"mentedu-backend/internal/model"
)

type ArticleCategoryRepositoryUseCase interface {
	Create(ctx context.Context, articleCategory *model.ArticleCategory) error
	Delete(ctx context.Context, articleID, categoryID int) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ArticleCategory, error)
	GetByArticleID(ctx context.Context, articleID int) ([]*model.ArticleCategory, error)
}

type ArticleCategoryRepository struct {
	db *gorm.DB
}

func NewArticleCategoryRepository(db *gorm.DB) ArticleCategoryRepositoryUseCase {
	return &ArticleCategoryRepository{db}
}

func (r *ArticleCategoryRepository) Create(ctx context.Context, articleCategory *model.ArticleCategory) error {
	if err := r.db.WithContext(ctx).Model(&model.ArticleCategory{}).Create(&articleCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *ArticleCategoryRepository) Delete(ctx context.Context, articleID, categoryID int) error {
	if err := r.db.WithContext(ctx).Model(&model.ArticleCategory{}).Where("article_id = ? AND category_id = ?", articleID, categoryID).Delete(&model.ArticleCategory{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ArticleCategoryRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ArticleCategory, error) {
	var articleCategories []*model.ArticleCategory

	q := repo.db.WithContext(ctx).Model(&model.ArticleCategory{})

	if query != "" {
		q = q.Where("name LIKE ?", "%"+query+"%")
	}

	if sort != "" && order != "" {
		q = q.Order(sort + " " + order)
	} else {
		q = q.Order("created_at DESC")
	}

	if limit > 0 {
		q = q.Limit(limit)
	}

	if offset > 0 {
		q = q.Offset(offset)
	}

	if err := q.Find(&articleCategories).Error; err != nil {
		return nil, err
	}

	return articleCategories, nil
}

func (r *ArticleCategoryRepository) GetByArticleID(ctx context.Context, articleID int) ([]*model.ArticleCategory, error) {
	var articleCategories []*model.ArticleCategory
	if err := r.db.WithContext(ctx).Model(&model.ArticleCategory{}).Where("article_id = ?", articleID).Find(&articleCategories).Error; err != nil {
		return nil, err
	}
	return articleCategories, nil
}
