package validator

import (
	"fmt"
	"net/http"
)

type ValidatorImpl struct {
	constraints []Constraint
}
type Validator interface {
	Validate(r *http.Request) error
}

func New(constraints []Constraint) *ValidatorImpl {
	return &ValidatorImpl{
		constraints: constraints,
	}
}

type Constraint int

const (
	User_Agent Constraint = iota
	Accept_Header
)

func (v *ValidatorImpl) Validate(r *http.Request) error {
	for _, c := range v.constraints {
		constraint := selectConstraint(c)
		if err := constraint(r); err != nil {
			return err
		}
	}
	return nil
}

func UserAgent(r *http.Request) error {
	// If the User-Agent header is missing, or the length of the header value is less than 40 bytes, return a HTTP 406 response.
	userAgent := r.UserAgent()
	if len(userAgent) == 0 {
		return fmt.Errorf("User-Agent header needs to contain a value that is larger than 40 bytes")
	}
	if len(userAgent) < 40 {
		return fmt.Errorf("User-Agent header needs to contain a value larger than 40 bytes, current value is of %d bytes, %s", len(userAgent), userAgent)
	}
	return nil
}

func AcceptHeader(r *http.Request) error {
	// If the Accept header is missing, return a HTTP 406 response.
	acceptHeader := r.Header.Get("Accept")
	if len(acceptHeader) == 0 {
		return fmt.Errorf("accept header needs to contain a value")
	}
	return nil
}

func selectConstraint(c Constraint) func(r *http.Request) error {
	switch c {
	case User_Agent:
		return func(r *http.Request) error {
			return UserAgent(r)
		}
	case Accept_Header:
		return func(r *http.Request) error {
			return AcceptHeader(r)
		}
	}
	return nil
}
