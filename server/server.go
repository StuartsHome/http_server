package server

import (
	"fmt"
	"http_server/validator"
	"net/http"
)

type ServerImpl struct {
	validator.Validator
}

var _ http.Handler = &ServerImpl{}

func New(validator validator.Validator) *ServerImpl {
	return &ServerImpl{
		Validator: validator,
	}
}

func (s *ServerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := s.Validator.Validate(r); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
