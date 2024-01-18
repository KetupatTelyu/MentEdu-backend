package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mentedu-backend/internal/model"
)

type ConsultationOrderRepositoryUseCase interface {
	Create(ctx context.Context, consultationOrder *model.ConsultationOrder) error
	Update(ctx context.Context, consultationOrder *model.ConsultationOrder) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationOrder, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.ConsultationOrder, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type ConsultationOrderRepository struct {
	db *gorm.DB
}

func NewConsultationOrderRepository(db *gorm.DB) ConsultationOrderRepositoryUseCase {
	return &ConsultationOrderRepository{db}
}

func (r *ConsultationOrderRepository) Create(ctx context.Context, consultationOrder *model.ConsultationOrder) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationOrder{}).Create(&consultationOrder).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultationOrderRepository) Update(ctx context.Context, consultationOrder *model.ConsultationOrder) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationOrder{}).Where("id = ?", consultationOrder.ID).Updates(consultationOrder).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsultationOrderRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationOrder{}).Delete(&model.ConsultationOrder{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ConsultationOrderRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationOrder, int64, error) {
	var consultationOrders []*model.ConsultationOrder

	q := repo.db.WithContext(ctx).Model(&model.ConsultationOrder{})

	uuid, err := uuid.Parse(query)

	if query != "" {
		if err == nil {
			q = q.Where("user_id = ?", uuid).Or("consultant_id = ?", uuid)
		} else {
			q = q.Where("purpose LIKE ?", "%"+query+"%")
		}
	}

	if sort != "" && order != "" {
		q = q.Order(sort + " " + order)
	} else {
		q = q.Order("created_at DESC")
	}

	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	if err := q.Find(&consultationOrders).Error; err != nil {
		return nil, 0, err
	}

	var total int64

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return consultationOrders, total, nil
}

func (r *ConsultationOrderRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.ConsultationOrder, error) {
	var consultationOrder model.ConsultationOrder

	if err := r.db.WithContext(ctx).Model(&model.ConsultationOrder{}).Where("id = ?", id).First(&consultationOrder).Error; err != nil {
		return nil, err
	}

	return &consultationOrder, nil
}

func (r *ConsultationOrderRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	if err := r.db.WithContext(ctx).Model(&model.ConsultationOrder{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
