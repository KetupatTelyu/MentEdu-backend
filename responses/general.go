package responses

type Response struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func NewResponse(data interface{}, status int, message string) *Response {
	return &Response{
		Data:    data,
		Status:  status,
		Message: message,
	}
}

type PaginatedResponse struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Total   int64       `json:"total"`
}

func NewPaginatedResponse(data interface{}, status int, message string, total int64) *PaginatedResponse {
	return &PaginatedResponse{
		Data:    data,
		Status:  status,
		Message: message,
		Total:   total,
	}
}
