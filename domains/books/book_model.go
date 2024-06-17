package books

type BookModel struct {
	ID          uint    `gorm:"column:id"`
	Code        string  `gorm:"column:code"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Price       float32 `gorm:"column:price"`
}

type Tabler interface {
	TableName() string
}

func (BookModel) TableName() string {
	return "books"
}
