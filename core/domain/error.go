package domain

import (
	"net/http"
)

type Error struct {
	Message string
	Code    int
	Error   string
}

func NewErr(msg string, code int) *Error {
	return &Error{
		Code:    code,
		Error:   http.StatusText(code),
		Message: msg,
	}
}
