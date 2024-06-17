package order

import (
	"time"
)

type OrderResponsePayload struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    OrderResponseData `json:"data"`
}
type OrderResponseData struct {
	OrderNo string `json:"order_no"`
}

type OrderListResponsePayload struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    []OrderListResponseData `json:"data"`
}
type OrderListResponseData struct {
	OrderNo       string                    `json:"order_no"`
	CreatedAt     *time.Time                `json:"created_at"`
	OrderStatus   string                    `json:"order_status"`
	PaymentStatus string                    `json:"payment_status"`
	OrderDetails  []OrderListDetailResponse `json:"order_details"`
}

type OrderListDetailResponse struct {
	BookTitle string `json:"book_title"`
	BookCode  string `json:"book_code"`
	Qty       int    `json:"qty"`
}
