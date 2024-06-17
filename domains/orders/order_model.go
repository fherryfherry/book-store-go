package orders

import "time"

type OrderModel struct {
	ID              uint       `gorm:"column:id"`
	CreatedAt       *time.Time `gorm:"column:created_at"`
	OrderNo         string     `gorm:"column:order_no"`
	CustomerID      uint       `gorm:"column:customer_id"`
	CustomerName    string     `gorm:"column:customer_name"`
	CustomerEmail   string     `gorm:"column:customer_email"`
	ShipmentAddress string     `gorm:"column:shipment_address"`
	GrandTotal      float32    `gorm:"column:grand_total"`
	OrderStatus     string     `gorm:"column:order_status"`
	PaymentStatus   string     `gorm:"column:payment_status"`
}

type Tabler interface {
	TableName() string
}

func (OrderModel) TableName() string {
	return "orders"
}
