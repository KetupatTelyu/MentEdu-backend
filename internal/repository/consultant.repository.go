package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mentedu-backend/internal/model"
)

type ConsultantRepositoryUseCase interface {
	Create(ctx context.Context, consultant *model.Consultant) error

	Update(ctx context.Context, consultant *model.Consultant) error

	Delete(ctx context.Context, id uuid.UUID) error

	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Consultant, int64, error)

	GetById(ctx context.Context, id uuid.UUID) (*model.Consultant, error)
}

type ConsultantRepository struct {
	db *gorm.DB
}

func NewConsultantRepository(db *gorm.DB) ConsultantRepositoryUseCase {
	return &ConsultantRepository{db}
}

func (r *ConsultantRepository) Create(ctx context.Context, consultant *model.Consultant) error {
	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Create(&consultant).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultantRepository) Update(ctx context.Context, consultant *model.Consultant) error {
	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Where("id = ?", consultant.ID).Updates(consultant).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Where("id = ?", id).Delete(&model.Consultant{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultantRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Consultant, int64, error) {
	var consultants []*model.Consultant
	var total int64

	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Where("name LIKE ?", "%"+query+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Where("name LIKE ?", "%"+query+"%").Order(sort + " " + order).Limit(limit).Offset(offset).Find(&consultants).Error; err != nil {
		return nil, 0, err
	}

	return consultants, total, nil
}

func (r *ConsultantRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Consultant, error) {
	var consultant model.Consultant

	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Where("id = ?", id).First(&consultant).Error; err != nil {
		return nil, err
	}

	return &consultant, nil
}
