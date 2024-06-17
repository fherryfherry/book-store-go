package registration

import (
	errCommon "booking-online/commons/error"
	"github.com/labstack/echo/v4"
)

func (r *Registration) RegisterHandler(c echo.Context) error {

	reqPayload := RegistrationRequestPayload{}
	if err := c.Bind(&reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	if err := c.Validate(reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	// Validate exists of account by email
	if r.customerSrv.CheckExistByEmail(reqPayload.Email) {
		return errCommon.ErrorResponseBadRequest(c, "Account already exists!")
	}

	// Insert process
	customerModel, err := r.customerSrv.CreateCustomer(reqPayload.FirstName, reqPayload.LastName, reqPayload.Email, reqPayload.Password)
	if err != nil {
		return errCommon.ErrorResponseInternalError(c, err.Error())
	}

	return c.JSON(200, RegistrationResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    RegistrationResponseData{customerModel.ID},
	})
}
