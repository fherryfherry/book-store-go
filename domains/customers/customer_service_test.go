package customers

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
		Logger: logger.Default.LogMode(logger.Info), // Set log level to Info to see detailed SQL queries
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

func TestFindByEmail(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	customerService := InitCustomerService(db)

	expectedCustomer := CustomerModel{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Password:  "hashedpassword",
	}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password"}).
		AddRow(expectedCustomer.ID, expectedCustomer.FirstName, expectedCustomer.LastName, expectedCustomer.Email, expectedCustomer.Password)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `customers` WHERE email = ? and deleted_at is null ORDER BY `customers`.`id` LIMIT ?")).
		WithArgs("johndoe@example.com", 1).
		WillReturnRows(rows)

	customer := customerService.FindByEmail("johndoe@example.com")

	if customer != expectedCustomer {
		t.Errorf("expected %+v, got %+v", expectedCustomer, customer)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCheckExistByEmail(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	customerService := InitCustomerService(db)

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password"}).
		AddRow(1, "John", "Doe", "johndoe@example.com", "hashedpassword")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `customers` WHERE email = ? and deleted_at is null ORDER BY `customers`.`id` LIMIT ?")).
		WithArgs("johndoe@example.com", 1).
		WillReturnRows(rows)

	exists := customerService.CheckExistByEmail("johndoe@example.com")

	if !exists {
		t.Errorf("expected true, got false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
