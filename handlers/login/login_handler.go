package login

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/jwt"
	passCommon "booking-online/commons/password"
	"github.com/labstack/echo/v4"
)

func (r *Login) LoginHandler(c echo.Context) error {

	reqPayload := LoginRequestPayload{}
	if err := c.Bind(&reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	if err := c.Validate(reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	customer := r.customerSrv.FindByEmail(reqPayload.Email)
	if customer.Email == "" {
		return errCommon.ErrorResponseBadRequest(c, "Email and or password is wrong!")
	}

	if !passCommon.CheckPasswordHash(reqPayload.Password, customer.Password) {
		return errCommon.ErrorResponseBadRequest(c, "Email and or password is wrong!")
	}

	token, err := jwt.CreateJwtToken(customer.ID, customer.FirstName, customer.Email)
	if err != nil {
		return errCommon.ErrorResponseInternalError(c, err.Error())
	}

	return c.JSON(200, LoginResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    LoginResponseData{token},
	})
}
