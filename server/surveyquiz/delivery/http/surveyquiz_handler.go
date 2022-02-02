package http

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	// case domain.ErrInternalServerError:
	// 	return http.StatusInternalServerError
	// case domain.ErrNotFound:
	// 	return http.StatusNotFound
	// case domain.ErrConflict:
	// 	return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
