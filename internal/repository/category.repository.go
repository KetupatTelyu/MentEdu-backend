package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mentedu-backend/internal/model"
)

type CategoryRepositoryUseCase interface {
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Category, error)
	GetById(ctx context.Context, id int) (*model.Category, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryUseCase {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Model(&model.Category{}).Create(&category).Error; err != nil {
		return errors.Wrap(err, "error creating category")
	}
	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Model(&model.Category{}).Where("id = ?", category.ID).Updates(category).Error; err != nil {
		return errors.Wrap(err, "error updating category")
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Model(&model.Category{}).Delete(&model.Category{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting category")
	}
	return nil
}

func (repo *CategoryRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Category, error) {
	var categories []*model.Category

	q := repo.db.WithContext(ctx).Model(&model.Category{})

	if query != "" {
		q = q.Where("name LIKE ?", "%"+query+"%")
	}

	if sort != "" && order != "" {
		q = q.Order(sort + " " + order)
	} else {
		q = q.Order("created_at DESC")
	}

	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	if err := q.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetById(ctx context.Context, id int) (*model.Category, error) {
	var category model.Category
	if err := r.db.WithContext(ctx).Model(&model.Category{}).First(&category, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting category by id")
	}
	return &category, nil
}
