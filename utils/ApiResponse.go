package utils

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func HandleSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Message: message,
		Data:    data,
	}
}

func HandleErrorResponse(message string, error string) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Error:   error,
	}
}
