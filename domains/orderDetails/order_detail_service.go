package orderDetails

import (
	"gorm.io/gorm"
)

func InitOrderDetailService(DbCon *gorm.DB) OrderDetailService {
	return OrderDetailService{DbCon: DbCon}
}

type OrderDetailService struct {
	DbCon *gorm.DB
}

func (c *OrderDetailService) GetListByOrderIds(ids []uint) map[uint][]OrderDetailModel {
	mapOrderList := map[uint][]OrderDetailModel{}
	var orderList []OrderDetailModel
	result := c.DbCon.Where("order_id IN ?", ids).Find(&orderList)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}

	for _, item := range orderList {
		mapOrderList[item.OrderID] = append(mapOrderList[item.OrderID], item)
	}

	return mapOrderList
}
