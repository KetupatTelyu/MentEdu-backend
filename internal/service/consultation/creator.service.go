package consultation

import (
	"context"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/repository"
)

type ConsultationCreatorUsecase interface {
	CreateConsultation(ctx context.Context, consultation *model.ConsultationOrder) error

	CreateConsultationDetail(ctx context.Context, consultationDetail *model.ConsultationDetail) error
}

type ConsultationCreator struct {
	consultationRepo       repository.ConsultationOrderRepositoryUseCase
	consultationDetailRepo repository.ConsultationDetailsRepositoryUseCase
}

func NewConsultationCreator(consultationRepo repository.ConsultationOrderRepositoryUseCase, consultationDetailRepo repository.ConsultationDetailsRepositoryUseCase) ConsultationCreatorUsecase {
	return &ConsultationCreator{
		consultationRepo:       consultationRepo,
		consultationDetailRepo: consultationDetailRepo,
	}
}

func (cc *ConsultationCreator) CreateConsultation(ctx context.Context, consultation *model.ConsultationOrder) error {
	if err := cc.consultationRepo.Create(ctx, consultation); err != nil {
		return err
	}

	return nil
}

func (cc *ConsultationCreator) CreateConsultationDetail(ctx context.Context, consultationDetail *model.ConsultationDetail) error {
	if err := cc.consultationDetailRepo.Create(ctx, consultationDetail); err != nil {
		return err
	}

	return nil
}
