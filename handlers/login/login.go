package login

import (
	"booking-online/domains/customers"
	"gorm.io/gorm"
)

type Login struct {
	customerSrv customers.CustomerService
}

func InitHandler(dbConnection *gorm.DB) Login {
	return Login{
		customerSrv: customers.InitCustomerService(dbConnection),
	}
}
