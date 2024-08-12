package validator

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAgent_OK(t *testing.T) {
	// Given.
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"
	validator := New([]Constraint{User_Agent})

	// When.
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("User-Agent", userAgent)

	// Then.
	require.NoError(t, validator.Validate(req))
}

func TestUserAgent_Error(t *testing.T) {
	// Given.
	userAgent := ""
	validator := New([]Constraint{User_Agent})

	// When.
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("User-Agent", userAgent)

	// Then.
	got := validator.Validate(req)
	assert.EqualError(t, got, "User-Agent header needs to contain a value that is larger than 40 bytes")
}

func TestUserAcceptHeader_OK(t *testing.T) {
	// Given.
	acceptHeader := "*/*"
	validator := New([]Constraint{Accept_Header})

	// When.
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("Accept", acceptHeader)

	// Then.
	require.NoError(t, validator.Validate(req))
}

func TestUserAcceptHeader_Error(t *testing.T) {
	// Given.
	acceptHeader := ""
	validator := New([]Constraint{Accept_Header})

	// When.
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("Accept", acceptHeader)

	// Then.
	got := validator.Validate(req)
	assert.EqualError(t, got, "accept header needs to contain a value")
}
