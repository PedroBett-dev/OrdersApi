package apperror

import "net/http"

var (
	ErrInvalidCpf          = New(http.StatusBadRequest, "cpf inválido")
	ErrInvalidCep          = New(http.StatusBadRequest, "cep inválido")
	ErrInvalidState        = New(http.StatusBadRequest, "uf inválida")
	ErrInvalidPhoneNumber  = New(http.StatusBadRequest, "telefone inválido")
	ErrInvalidEmail        = New(http.StatusBadRequest, "email inválido")
	ErrClientAlreadyExists = New(http.StatusConflict, "cliente já cadastrado")
	ErrClientNotFound      = New(http.StatusNotFound, "Cliente nao cadastrado")
)
