package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	id := uint(1)
	name := "John Doe"
	email := "john.doe@example.com"

	token, err := CreateJwtToken(id, name, email)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123456"), nil
	})

	claims, ok := parsedToken.Claims.(*JwtCustomClaims)
	assert.True(t, ok)
	assert.Equal(t, id, claims.ID)
	assert.Equal(t, name, claims.Name)
	assert.Equal(t, email, claims.Email)
}

func TestGetClaim(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token, _ := CreateJwtToken(1, "John Doe", "john.doe@example.com")

	parsedToken, _ := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret123456"), nil
	})

	c.Set("user", parsedToken)

	claims := GetClaim(c)
	assert.Equal(t, uint(1), claims.ID)
	assert.Equal(t, "John Doe", claims.Name)
	assert.Equal(t, "john.doe@example.com", claims.Email)
}

func TestInitMiddlewareJwt(t *testing.T) {
	e := echo.New()
	middleware := InitMiddlewareJwt()
	e.Use(middleware)

	token, _ := CreateJwtToken(1, "John Doe", "john.doe@example.com")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		claims := GetClaim(c)
		return c.JSON(http.StatusOK, claims)
	}

	e.GET("/", handler)

	if assert.NoError(t, middleware(handler)(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := `{"name":"John Doe","id":1,"email":"john.doe@example.com","exp":`
		assert.Contains(t, rec.Body.String(), expected)
	}
}
