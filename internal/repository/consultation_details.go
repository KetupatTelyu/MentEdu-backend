package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mentedu-backend/internal/model"
)

type ConsultationDetailsRepositoryUseCase interface {
	Create(ctx context.Context, consultationDetails *model.ConsultationDetail) error
	Update(ctx context.Context, consultationDetails *model.ConsultationDetail) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationDetail, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.ConsultationDetail, error)
}

type ConsultationDetailsRepository struct {
	db *gorm.DB
}

func NewConsultationDetailsRepository(db *gorm.DB) ConsultationDetailsRepositoryUseCase {
	return &ConsultationDetailsRepository{db}
}

func (r *ConsultationDetailsRepository) Create(ctx context.Context, consultationDetails *model.ConsultationDetail) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationDetail{}).Create(&consultationDetails).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultationDetailsRepository) Update(ctx context.Context, consultationDetails *model.ConsultationDetail) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationDetail{}).Where("id = ?", consultationDetails.ID).Updates(consultationDetails).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultationDetailsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationDetail{}).Delete(&model.ConsultationDetail{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ConsultationDetailsRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationDetail, int64, error) {
	var consultationDetails []*model.ConsultationDetail

	q := repo.db.WithContext(ctx).Model(&model.ConsultationDetail{})

	if query != "" {
		q = q.Where("name ILIKE ?", "%"+query+"%")
	}

	if sort != "" && order != "" {
		q = q.Order(sort + " " + order)
	} else {
		q = q.Order("created_at DESC")
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Limit(limit).Offset(offset).Find(&consultationDetails).Error; err != nil {
		return nil, 0, err
	}

	return consultationDetails, total, nil
}

func (r *ConsultationDetailsRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.ConsultationDetail, error) {
	var consultationDetail model.ConsultationDetail

	if err := r.db.WithContext(ctx).Model(&model.ConsultationDetail{}).Where("id = ?", id).First(&consultationDetail).Error; err != nil {
		return nil, err
	}

	return &consultationDetail, nil
}
