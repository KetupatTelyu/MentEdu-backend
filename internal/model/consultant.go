package model

import "github.com/google/uuid"

// extends user model
type Consultant struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	UserID        uuid.UUID `json:"user_id" gorm:"type:uuid"`
	IsAvailable   bool      `json:"is_available"`
	Occupation    string    `json:"occupation"`
	Company       string    `json:"company"`
	Sessions      int       `json:"sessions"`
	ReviewsTotal  int       `json:"reviews"`
	AvgAttendance float64   `json:"avg_attendance"`
	User          User      `json:"user" gorm:"foreignKey:ID"`
	Auditable
}

func NewConsultant(
	id uuid.UUID,
	userID uuid.UUID,
	isAvailable bool,
	occupation string,
	company string,
	sessions int,
	reviewsTotal int,
	avgAttendance float64,
	createdBy string,
) *Consultant {
	return &Consultant{
		ID:            id,
		UserID:        userID,
		IsAvailable:   isAvailable,
		Occupation:    occupation,
		Company:       company,
		Sessions:      sessions,
		ReviewsTotal:  reviewsTotal,
		AvgAttendance: avgAttendance,
		Auditable:     NewAuditable(createdBy),
	}
}

func (model *Consultant) MapUpdateFrom(from *Consultant) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"user_id":        model.UserID,
			"is_available":   model.IsAvailable,
			"occupation":     model.Occupation,
			"company":        model.Company,
			"sessions":       model.Sessions,
			"reviews_total":  model.ReviewsTotal,
			"avg_attendance": model.AvgAttendance,
		}
	}

	m := map[string]interface{}{}

	if from.UserID != model.UserID {
		m["user_id"] = model.UserID
	}

	if from.IsAvailable != model.IsAvailable {
		m["is_available"] = model.IsAvailable
	}

	if from.Occupation != model.Occupation {
		m["occupation"] = model.Occupation
	}

	if from.Company != model.Company {
		m["company"] = model.Company
	}

	if from.Sessions != model.Sessions {
		m["sessions"] = model.Sessions
	}

	if from.ReviewsTotal != model.ReviewsTotal {
		m["reviews_total"] = model.ReviewsTotal
	}

	if from.AvgAttendance != model.AvgAttendance {
		m["avg_attendance"] = model.AvgAttendance
	}

	if len(m) == 0 {
		return nil
	}

	return &m
}
