package order

import (
	"booking-online/domains/orderDetails"
	"booking-online/domains/orders"
	"gorm.io/gorm"
)

type Order struct {
	orderSrv       orders.OrderService
	orderDetailSrv orderDetails.OrderDetailService
}

func InitHandler(dbConnection *gorm.DB) Order {
	return Order{
		orderSrv:       orders.InitOrderService(dbConnection),
		orderDetailSrv: orderDetails.InitOrderDetailService(dbConnection),
	}
}
