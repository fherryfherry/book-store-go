package book

import (
	"booking-online/domains/books"
	"gorm.io/gorm"
)

type Book struct {
	bookSrv books.BookService
}

func InitHandler(dbConnection *gorm.DB) Book {
	return Book{
		bookSrv: books.InitBookService(dbConnection),
	}
}
