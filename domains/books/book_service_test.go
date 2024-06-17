package books

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql db, got error: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, config)
	if err != nil {
		t.Fatalf("failed to open gorm db, got error: %v", err)
	}

	return db, mock
}

func TestFindByCode(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	bookService := InitBookService(db)

	// Define the expected book model
	expectedBook := BookModel{
		ID:          1,
		Code:        "ABC123",
		Title:       "Test Book",
		Description: "A book for testing",
		Price:       19.99,
	}

	rows := sqlmock.NewRows([]string{"id", "code", "title", "description", "price"}).
		AddRow(expectedBook.ID, expectedBook.Code, expectedBook.Title, expectedBook.Description, expectedBook.Price)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE code = ? and deleted_at is null ORDER BY `books`.`id` LIMIT ?")).
		WithArgs("ABC123", 1).
		WillReturnRows(rows)

	book := bookService.FindByCode("ABC123")
	if book != expectedBook {
		t.Errorf("expected %v, got %v", expectedBook, book)
	}
}

func TestGetList(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	bookService := InitBookService(db)

	// Define the expected book models
	expectedBooks := []BookModel{
		{ID: 1, Code: "ABC123", Title: "Test Book 1", Description: "A book for testing 1", Price: 19.99},
		{ID: 2, Code: "DEF456", Title: "Test Book 2", Description: "A book for testing 2", Price: 29.99},
	}

	rows := sqlmock.NewRows([]string{"id", "code", "title", "description", "price"}).
		AddRow(expectedBooks[0].ID, expectedBooks[0].Code, expectedBooks[0].Title, expectedBooks[0].Description, expectedBooks[0].Price).
		AddRow(expectedBooks[1].ID, expectedBooks[1].Code, expectedBooks[1].Title, expectedBooks[1].Description, expectedBooks[1].Price)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE deleted_at is null")).
		WillReturnRows(rows)

	books := bookService.GetList()
	if len(books) != len(expectedBooks) {
		t.Errorf("expected %d books, got %d", len(expectedBooks), len(books))
	}

	for i, book := range books {
		if book != expectedBooks[i] {
			t.Errorf("expected %v, got %v", expectedBooks[i], book)
		}
	}
}
