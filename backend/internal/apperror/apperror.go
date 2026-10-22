package apperror

type AppError struct {
	Status  int
	Message string
}

func New(status int, message string) *AppError {
	return &AppError{Status: status, Message: message}
}

func (e *AppError) Error() string {
	return e.Message
}
