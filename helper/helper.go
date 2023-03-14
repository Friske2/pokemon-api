package helper

import (
	"log"
	"net/http"

	"github.com/Friske2/pokemon-api/domain"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Fatal(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	case domain.ErrBadRequestBody:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
