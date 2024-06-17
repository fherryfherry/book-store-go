package validator

import (
	"github.com/go-playground/validator"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=130"`
}

func TestCustomValidator_Validate(t *testing.T) {
	v := validator.New()
	cv := CustomValidator{Validator: v}

	tests := []struct {
		name      string
		input     interface{}
		wantError bool
	}{
		{
			name: "Valid struct",
			input: TestStruct{
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Age:   30,
			},
			wantError: false,
		},
		{
			name: "Missing name",
			input: TestStruct{
				Email: "john.doe@example.com",
				Age:   30,
			},
			wantError: true,
		},
		{
			name: "Invalid email",
			input: TestStruct{
				Name:  "John Doe",
				Email: "invalid-email",
				Age:   30,
			},
			wantError: true,
		},
		{
			name: "Age out of range",
			input: TestStruct{
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Age:   150,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cv.Validate(tt.input)
			if tt.wantError {
				assert.Error(t, err)
				httpError, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, http.StatusBadRequest, httpError.Code)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
