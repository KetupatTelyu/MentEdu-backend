package resource

type ConsultationOrderRequest struct {
	UserID       string `json:"user_id" binding:"required"`
	Purpose      string `json:"purpose" binding:"required"`
	DateTime     string `json:"date_time" binding:"required" time_format:"2006-01-02 15:04:05"`
	Status       string `json:"status" binding:"required"`
	ConsultantID string `json:"consultant_id" binding:"required"`
}

type ConsultationDetailRequest struct {
	ConsultationOrderID string `json:"consultation_order_id" binding:"required"`
	MeetingURL          string `json:"meeting_url" binding:"required"`
}

type ConsultationStatusRequest struct {
	ConsultationOrderID string `json:"consultation_order_id" binding:"required"`
	Status              string `json:"status" binding:"required"`
}
