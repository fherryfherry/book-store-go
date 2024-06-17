package order

import (
	errCommon "booking-online/commons/error"
	"booking-online/commons/jwt"
	"booking-online/domains/orderDetails"
	"booking-online/domains/orders"
	"github.com/labstack/echo/v4"
)

func (r *Order) CheckoutOrder(c echo.Context) error {

	reqPayload := OrderRequestPayload{}
	if err := c.Bind(&reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	if err := c.Validate(reqPayload); err != nil {
		return errCommon.ErrorResponseBadRequest(c, err.Error())
	}

	// Get jwt account
	customer := jwt.GetClaim(c)

	// Process checkout order
	var bookList []orders.CreateOrderBook
	for _, pb := range reqPayload.BookList {
		bookList = append(bookList, orders.CreateOrderBook{
			BookCode: pb.BookCode,
			Qty:      pb.Qty,
		})
	}

	payload := orders.CreateOrderPayload{
		CustomerID:      customer.ID,
		CustomerName:    customer.Name,
		CustomerEmail:   customer.Email,
		ShipmentAddress: reqPayload.ShipmentAddress,
		BookList:        bookList,
	}
	orderNo, err := r.orderSrv.CreateOrder(payload)
	if err != nil {
		return errCommon.ErrorResponseInternalError(c, err.Error())
	}

	return c.JSON(200, OrderResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    OrderResponseData{orderNo},
	})
}

func (r *Order) MyOrder(c echo.Context) error {

	// Get jwt account
	customer := jwt.GetClaim(c)

	result := r.orderSrv.GetListByCustomer(customer.ID)
	orderIds := extractOrderIds(result)
	orderDetailMaps := r.orderDetailSrv.GetListByOrderIds(orderIds)
	resultOrderList := []OrderListResponseData{}
	for _, item := range result {
		resultOrderList = append(resultOrderList, OrderListResponseData{
			OrderNo:       item.OrderNo,
			CreatedAt:     item.CreatedAt,
			OrderStatus:   item.OrderStatus,
			PaymentStatus: item.PaymentStatus,
			OrderDetails:  toOrderListDetailResponse(orderDetailMaps[item.ID]),
		})
	}

	return c.JSON(200, OrderListResponsePayload{
		Status:  200,
		Message: "SUCCESS",
		Data:    resultOrderList,
	})
}

func toOrderListDetailResponse(orderDetailList []orderDetails.OrderDetailModel) []OrderListDetailResponse {
	var orderListDetailResp []OrderListDetailResponse
	for _, o := range orderDetailList {
		orderListDetailResp = append(orderListDetailResp, OrderListDetailResponse{
			BookTitle: o.BookTitle,
			BookCode:  o.BookCode,
			Qty:       o.Qty,
		})
	}
	return orderListDetailResp
}

func extractOrderIds(orderList []orders.OrderModel) []uint {
	orderIds := []uint{}
	for _, item := range orderList {
		orderIds = append(orderIds, item.ID)
	}
	return orderIds
}
