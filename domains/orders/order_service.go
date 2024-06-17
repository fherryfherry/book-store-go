package orders

import (
	"booking-online/domains/books"
	orderDetails "booking-online/domains/orderDetails"
	"fmt"
	"gorm.io/gorm"
)

func InitOrderService(DbCon *gorm.DB) OrderService {
	return OrderService{DbCon: DbCon}
}

type OrderService struct {
	DbCon *gorm.DB
}

func (c *OrderService) GetListByCustomer(customerId uint) []OrderModel {
	var orderList []OrderModel
	result := c.DbCon.Order("id desc").Find(&orderList, "customer_id = ? and deleted_at is null", customerId)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}
	return orderList
}

func (c *OrderService) CountOrder() int64 {
	var total int64
	c.DbCon.Table("orders").Count(&total)
	return total
}

func (c *OrderService) CreateOrder(payload CreateOrderPayload) (string, error) {
	total := c.CountOrder() + 1
	orderNo := "BS" + fmt.Sprintf("%03d", total)
	err := c.DbCon.Transaction(func(tx *gorm.DB) error {
		orderMaster := new(OrderModel)
		orderMaster.OrderNo = orderNo
		orderMaster.OrderStatus = "COMPLETED"
		orderMaster.PaymentStatus = "PAID"
		orderMaster.CustomerID = payload.CustomerID
		orderMaster.CustomerName = payload.CustomerName
		orderMaster.CustomerEmail = payload.CustomerEmail
		orderMaster.ShipmentAddress = payload.ShipmentAddress
		orderCreateResult := tx.Create(&orderMaster)
		if orderCreateResult.Error != nil {
			return orderCreateResult.Error
		}

		var grandTotal float32
		var details []*orderDetails.OrderDetailModel
		for _, item := range payload.BookList {
			var bookMdl books.BookModel
			bookResult := c.DbCon.Table("books").Where("code = ?", item.BookCode).First(&bookMdl)
			if bookResult.Error != nil {
				return bookResult.Error
			}
			totalPrice := bookMdl.Price * float32(item.Qty)
			details = append(details, &orderDetails.OrderDetailModel{
				OrderID:    orderMaster.ID,
				BookCode:   item.BookCode,
				BookTitle:  bookMdl.Title,
				BookPrice:  bookMdl.Price,
				Qty:        item.Qty,
				TotalPrice: totalPrice,
			})
			grandTotal += totalPrice
		}
		detailResult := tx.Create(details)
		if detailResult.Error != nil {
			return detailResult.Error
		}

		// Re-Update Order
		orderMaster.GrandTotal = grandTotal
		tx.Updates(&orderMaster)

		return nil
	})

	if err != nil {
		return "", err
	}

	return orderNo, nil
}
