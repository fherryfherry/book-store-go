package customers

import (
	passCommon "booking-online/commons/password"
	"gorm.io/gorm"
)

func InitCustomerService(DbCon *gorm.DB) CustomerService {
	return CustomerService{DbCon: DbCon}
}

type CustomerService struct {
	DbCon *gorm.DB
}

func (c *CustomerService) FindByEmail(email string) CustomerModel {
	customerData := CustomerModel{}
	result := c.DbCon.First(&customerData, "email = ? and deleted_at is null", email)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return CustomerModel{}
	}
	return customerData
}

func (c *CustomerService) CheckExistByEmail(email string) bool {
	customerData := CustomerModel{}
	result := c.DbCon.First(&customerData, "email = ? and deleted_at is null", email)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return false
	}
	if customerData.Email != "" {
		return true
	} else {
		return false
	}
}

func (c *CustomerService) CreateCustomer(firstName string, lastName string, email string, password string) (*CustomerModel, error) {
	newData := new(CustomerModel)
	newData.FirstName = firstName
	newData.LastName = lastName

	newData.Email = email
	passHashed, err := passCommon.HashPassword(password)
	if err != nil {
		return nil, err
	}
	newData.Password = passHashed

	result := c.DbCon.Create(&newData)
	if result.Error != nil {
		return nil, result.Error
	}

	return newData, nil
}
