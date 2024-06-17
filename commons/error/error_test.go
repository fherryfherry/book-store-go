package error

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorResponseBadRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	message := "Bad Request Error"
	err := ErrorResponseBadRequest(c, message)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		expected := `{"code":400,"message":"` + message + `"}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}

func TestErrorResponseInternalError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	message := "Internal Server Error"
	err := ErrorResponseInternalError(c, message)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		expected := `{"code":500,"message":"` + message + `"}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
