package order

type OrderRequestPayload struct {
	ShipmentAddress string             `json:"shipment_address" validate:"required"`
	BookList        []OrderRequestBook `json:"book_list" validate:"required"`
}
type OrderRequestBook struct {
	BookCode string `json:"book_code"`
	Qty      int    `json:"qty"`
}
