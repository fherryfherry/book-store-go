package orderDetails

type OrderDetailModel struct {
	OrderID    uint    `gorm:"column:order_id"`
	BookCode   string  `gorm:"column:book_code"`
	BookTitle  string  `gorm:"column:book_title"`
	BookPrice  float32 `gorm:"column:book_price"`
	Qty        int     `gorm:"column:qty"`
	TotalPrice float32 `gorm:"column:total_price"`
}

type Tabler interface {
	TableName() string
}

func (OrderDetailModel) TableName() string {
	return "order_details"
}
