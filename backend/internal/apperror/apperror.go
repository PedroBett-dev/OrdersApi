package apperror

import "net/http"

type AppError struct {
	Status  int
	Message string
}

/*--------------------------------------------------------*/
func New(status int, message string) *AppError {
	return &AppError{Status: status, Message: message}
}

/*--------------------------------------------------------*/
func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrInvalidCpf          = New(http.StatusBadRequest, "cpf inválido")
	ErrClientAlreadyExists = New(http.StatusConflict, "cliente já cadastrado")
)
