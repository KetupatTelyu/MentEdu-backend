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
