package registration

import (
	"booking-online/domains/customers"
	"gorm.io/gorm"
)

type Registration struct {
	customerSrv customers.CustomerService
}

func InitHandler(dbConnection *gorm.DB) Registration {
	return Registration{
		customerSrv: customers.InitCustomerService(dbConnection),
	}
}
