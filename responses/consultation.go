package responses

type ConsultationOrderResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Purpose      string `json:"purpose"`
	DateTime     string `json:"date_time"`
	Status       string `json:"status"`
	ConsultantID string `json:"consultant_id"`
}

type ConsultationDetailResponse struct {
	ID                string                    `json:"id"`
	ConsultationID    string                    `json:"consultation_id"`
	MeetingURL        string                    `json:"meeting_url"`
	ConsultationOrder ConsultationOrderResponse `json:"consultation_order"`
	Consultant        ProfileResponse           `json:"consultant"`
	Consultee         ProfileResponse           `json:"consultee"`
}
