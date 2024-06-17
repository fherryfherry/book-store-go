package customers

type CustomerModel struct {
	ID        uint   `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (CustomerModel) TableName() string {
	return "customers"
}
