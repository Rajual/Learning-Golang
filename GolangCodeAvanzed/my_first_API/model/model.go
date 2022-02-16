package model

import "errors"

var (
	//ErrPersonCanNotNil la persona no pude ser nula.
	ErrPersonCanNotNil = errors.New("la persona no puede ser nula")
	//ErrIDPersonDoesNotExists la perosna no exite
	ErrIDPersonDoesNotExists = errors.New("la persona no existe")
)
