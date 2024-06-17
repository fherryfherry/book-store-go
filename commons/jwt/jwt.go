package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JwtToken struct {
	Ctx echo.Context
}

func CreateJwtToken(id uint, name string, email string) (string, error) {
	claims := &JwtCustomClaims{
		name,
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetClaim(ctx echo.Context) *JwtCustomClaims {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func InitMiddlewareJwt() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(viper.GetString("jwt.secret")),
	}
	return echojwt.WithConfig(config)
}
