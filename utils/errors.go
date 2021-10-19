package utils

import "net/http"

type ICustomError interface {
	Error() string
	Code() int
}

type customError struct {
	err  error
	code int
}

func (c customError) Error() string {
	return c.err.Error()
}

func (c customError) Code() int {
	return http.StatusBadRequest
}

func Error(err error) ICustomError {
	return ServerError(err)
}

func ClientError(err error) ICustomError {
	return customError{err, http.StatusBadRequest}
}

func ServerError(err error) ICustomError {
	return customError{err, http.StatusInternalServerError}
}
