package model

import "github.com/google/uuid"

var (
	consultationDetailTableName = "consultation_details"
)

type ConsultationDetail struct {
	ID                uuid.UUID         `json:"id"`
	ConsultationID    uuid.UUID         `json:"consultation_id"`
	MeetingURL        string            `json:"meeting_url"`
	ConsultationOrder ConsultationOrder `json:"consultation_order" gorm:"foreignKey:ConsultationID"`
	Auditable
}

func NewConsultationDetail(
	id uuid.UUID,
	consultationID uuid.UUID,
	meetingURL string,
	createdBy string,
) *ConsultationDetail {
	return &ConsultationDetail{
		ID:             id,
		ConsultationID: consultationID,
		MeetingURL:     meetingURL,
		Auditable:      NewAuditable(createdBy),
	}
}

func (*ConsultationDetail) TableName() string {
	return consultationDetailTableName
}

func (model *ConsultationDetail) MapUpdateFrom(from *ConsultationDetail) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"consultation_id": model.ConsultationID,
			"meeting_url":     model.MeetingURL,
		}
	}

	m := map[string]interface{}{}

	if from.ConsultationID != model.ConsultationID {
		m["consultation_id"] = model.ConsultationID
	}

	if from.MeetingURL != model.MeetingURL {
		m["meeting_url"] = model.MeetingURL
	}

	if len(m) == 0 {
		return nil
	}

	return &m
}
