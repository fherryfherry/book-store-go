package book

type BookResponsePayload struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    []BookResponsePayloadData `json:"data"`
}
type BookResponsePayloadData struct {
	Code  string  `json:"code"`
	Title string  `json:"title"`
	Desc  string  `json:"description"`
	Price float32 `json:"price"`
}
