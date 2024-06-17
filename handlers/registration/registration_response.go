package registration

type RegistrationResponsePayload struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    RegistrationResponseData `json:"data"`
}
type RegistrationResponseData struct {
	ID uint `json:"id"`
}
