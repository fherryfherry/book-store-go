package login

type LoginResponsePayload struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}
type LoginResponseData struct {
	Token string `json:"token"`
}
