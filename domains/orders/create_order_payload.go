package orders

type CreateOrderPayload struct {
	CustomerID      uint
	CustomerName    string
	CustomerEmail   string
	ShipmentAddress string
	BookList        []CreateOrderBook
}

type CreateOrderBook struct {
	BookCode string `json:"book_code"`
	Qty      int    `json:"qty"`
}
