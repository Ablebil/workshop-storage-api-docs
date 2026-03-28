package apierror

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

func New(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}
