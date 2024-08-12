package server

import (
	"fmt"
	"http_server/validator"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServeHTTP_OK(t *testing.T) {
	// Given.
	acceptHeader := "*/*"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("User-Agent", userAgent)

	// When.
	response := httptest.NewRecorder()

	constraints := []validator.Constraint{validator.User_Agent, validator.Accept_Header}
	validator := validator.New(constraints)
	s := New(validator)
	s.ServeHTTP(response, req)

	// Then.
	got := response.Result().StatusCode
	want := http.StatusOK
	assert.Equal(t, want, got)
}

func TestServeHTTP_UserAgent406(t *testing.T) {
	// Given.
	userAgent := "curl/8.7.1"
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("User-Agent", userAgent)

	// When.
	response := httptest.NewRecorder()
	constraints := []validator.Constraint{validator.User_Agent}
	validator := validator.New(constraints)
	s := New(validator)
	s.ServeHTTP(response, req)

	// Then.
	got := response.Result().StatusCode
	gotBody := response.Body.String()
	want := http.StatusNotAcceptable
	assert.Equal(t, want, got)
	assert.Equal(t, "User-Agent header needs to contain a value larger than 40 bytes, current value is of 10 bytes, curl/8.7.1", gotBody)
}

func TestServeHTTP_UserAgent406_Empty(t *testing.T) {
	// Given.
	acceptHeader := "*/*"
	userAgent := ""
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("User-Agent", userAgent)

	// When.
	response := httptest.NewRecorder()
	constraints := []validator.Constraint{validator.User_Agent}
	validator := validator.New(constraints)
	s := New(validator)
	s.ServeHTTP(response, req)

	// Then.
	got := response.Result().StatusCode
	gotBody := response.Body.String()
	want := http.StatusNotAcceptable
	assert.Equal(t, want, got)
	assert.Equal(t, "User-Agent header needs to contain a value that is larger than 40 bytes", gotBody)
}

func TestServeHTTP_Accept406(t *testing.T) {
	// Given.
	acceptHeader := ""
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"
	req, err := http.NewRequest(http.MethodGet, "", nil)
	require.NoError(t, err)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("User-Agent", userAgent)

	// When.
	response := httptest.NewRecorder()

	constraints := []validator.Constraint{validator.Accept_Header}
	validator := validator.New(constraints)
	s := New(validator)
	s.ServeHTTP(response, req)

	// Then.
	got := response.Result().StatusCode
	gotBody := response.Body.String()
	want := http.StatusNotAcceptable
	assert.Equal(t, want, got)
	assert.Equal(t, "accept header needs to contain a value", gotBody)
}

func TestBytes(t *testing.T) {
	const nihongo = "日本語"
	t.Log(len(nihongo))
	t.Log(len(fmt.Sprintf("%x", nihongo)))
	t.Logf("% x", nihongo)
}
