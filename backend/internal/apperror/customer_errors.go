package apperror

import "net/http"

var (
	ErrInvalidCpf            = New(http.StatusBadRequest, "cpf inválido")
	ErrInvalidCep            = New(http.StatusBadRequest, "cep inválido")
	ErrInvalidState          = New(http.StatusBadRequest, "uf inválida")
	ErrInvalidPhoneNumber    = New(http.StatusBadRequest, "telefone inválido")
	ErrInvalidEmail          = New(http.StatusBadRequest, "email inválido")
	ErrCustomerAlreadyExists = New(http.StatusConflict, "customere já cadastrado")
	ErrCustomerNotFound      = New(http.StatusNotFound, "Customere nao cadastrado")
)
