package consultation

import (
	"context"
	"github.com/google/uuid"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/repository"
)

type ConsultationFinderUsecase interface {
	FindConsultation(ctx context.Context, id uuid.UUID) (*model.ConsultationOrder, error)

	FindConsultationDetail(ctx context.Context, id uuid.UUID) (*model.ConsultationDetail, error)

	FindAllConsultation(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationOrder, int64, error)

	FindAllConsultationDetail(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationDetail, int64, error)
}

type ConsultationFinder struct {
	consultationRepo       repository.ConsultationOrderRepositoryUseCase
	consultationDetailRepo repository.ConsultationDetailsRepositoryUseCase
}

func NewConsultationFinder(consultationRepo repository.ConsultationOrderRepositoryUseCase, consultationDetailRepo repository.ConsultationDetailsRepositoryUseCase) ConsultationFinderUsecase {
	return &ConsultationFinder{
		consultationRepo:       consultationRepo,
		consultationDetailRepo: consultationDetailRepo,
	}
}

func (cc *ConsultationFinder) FindConsultation(ctx context.Context, id uuid.UUID) (*model.ConsultationOrder, error) {
	consultation, err := cc.consultationRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return consultation, nil
}

func (cc *ConsultationFinder) FindConsultationDetail(ctx context.Context, id uuid.UUID) (*model.ConsultationDetail, error) {
	consultationDetail, err := cc.consultationDetailRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return consultationDetail, nil
}

func (cc *ConsultationFinder) FindAllConsultation(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationOrder, int64, error) {
	consultations, total, err := cc.consultationRepo.GetAll(ctx, query, sort, order, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return consultations, total, nil
}

func (cc *ConsultationFinder) FindAllConsultationDetail(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.ConsultationDetail, int64, error) {
	consultationDetails, total, err := cc.consultationDetailRepo.GetAll(ctx, query, sort, order, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return consultationDetails, total, nil
}
