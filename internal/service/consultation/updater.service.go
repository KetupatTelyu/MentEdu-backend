package consultation

import (
	"context"
	"github.com/google/uuid"
	"mentedu-backend/internal/repository"
)

type ConsultationUpdaterUsecase interface {
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type ConsultationUpdater struct {
	consultationRepo repository.ConsultationOrderRepositoryUseCase
}

func NewConsultationUpdater(consultationRepo repository.ConsultationOrderRepositoryUseCase) ConsultationUpdaterUsecase {
	return &ConsultationUpdater{
		consultationRepo: consultationRepo,
	}
}

func (cc *ConsultationUpdater) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	if err := cc.consultationRepo.UpdateStatus(ctx, id, status); err != nil {
		return err
	}

	return nil
}
