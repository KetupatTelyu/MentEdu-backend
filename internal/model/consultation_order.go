package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	consultationOrderTableName = "consultation_orders"
)

const (
	ConsultationOrderStatusPending  = "pending"
	ConsultationOrderStatusAccepted = "accepted"
	ConsultationOrderStatusRejected = "rejected"
	ConsultationOrderStatusCanceled = "canceled"
	ConsultationOrderStatusDone     = "done"
)

type ConsultationOrder struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Purpose      string    `json:"purpose"`
	DateTime     time.Time `json:"date_time"`
	Status       string    `json:"status"`
	ConsultantID uuid.UUID `json:"consultant_id,omitempty" gorm:"type:uuid"`
	Consultant   User      `json:"consultant" gorm:"foreignKey:ConsultantID"`
	Consultee    User      `json:"consultee" gorm:"foreignKey:UserID"`
	Auditable
}

func NewConsultationOrder(
	id uuid.UUID,
	userID uuid.UUID,
	purpose string,
	dateTime time.Time,
	status string,
	consultantID uuid.UUID,
	createdBy string,
) *ConsultationOrder {
	return &ConsultationOrder{
		ID:           id,
		UserID:       userID,
		Purpose:      purpose,
		DateTime:     dateTime,
		Status:       status,
		ConsultantID: consultantID,
		Auditable:    NewAuditable(createdBy),
	}
}

func (*ConsultationOrder) TableName() string {
	return consultationOrderTableName
}

func (model *ConsultationOrder) MapUpdateFrom(from *ConsultationOrder) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"user_id":       model.UserID,
			"purpose":       model.Purpose,
			"date_time":     model.DateTime,
			"status":        model.Status,
			"consultant_id": model.ConsultantID,
		}
	}

	return &map[string]interface{}{
		"user_id":       from.UserID,
		"purpose":       from.Purpose,
		"date_time":     from.DateTime,
		"status":        from.Status,
		"consultant_id": from.ConsultantID,
	}
}
