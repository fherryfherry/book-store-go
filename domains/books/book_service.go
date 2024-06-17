package books

import (
	"gorm.io/gorm"
)

func InitBookService(DbCon *gorm.DB) BookService {
	return BookService{DbCon: DbCon}
}

type BookService struct {
	DbCon *gorm.DB
}

func (c *BookService) FindByCode(code string) BookModel {
	bookModel := BookModel{}
	result := c.DbCon.First(&bookModel, "code = ? and deleted_at is null", code)
	if result.Error != nil && result.Error.Error() != "record not found" {
		return BookModel{}
	}
	return bookModel
}

func (c *BookService) GetList() []BookModel {
	bookModels := []BookModel{}
	result := c.DbCon.Find(&bookModels, "deleted_at is null")
	if result.Error != nil && result.Error.Error() != "record not found" {
		return nil
	}
	return bookModels
}
