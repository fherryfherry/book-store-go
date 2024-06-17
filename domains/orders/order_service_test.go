package orders

import (
	"regexp"
	"testing"
	"time"

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

func TestGetListByCustomer(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	orderService := InitOrderService(db)

	customerID := uint(1)
	orderTime := time.Now()
	expectedOrders := []OrderModel{
		{ID: 1, CreatedAt: &orderTime, OrderNo: "BS001", CustomerID: customerID, CustomerName: "John Doe", CustomerEmail: "johndoe@example.com", ShipmentAddress: "123 Test St", GrandTotal: 100.0, OrderStatus: "COMPLETED", PaymentStatus: "PAID"},
		{ID: 2, CreatedAt: &orderTime, OrderNo: "BS002", CustomerID: customerID, CustomerName: "John Doe", CustomerEmail: "johndoe@example.com", ShipmentAddress: "123 Test St", GrandTotal: 150.0, OrderStatus: "COMPLETED", PaymentStatus: "PAID"},
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "order_no", "customer_id", "customer_name", "customer_email", "shipment_address", "grand_total", "order_status", "payment_status"}).
		AddRow(expectedOrders[0].ID, expectedOrders[0].CreatedAt, expectedOrders[0].OrderNo, expectedOrders[0].CustomerID, expectedOrders[0].CustomerName, expectedOrders[0].CustomerEmail, expectedOrders[0].ShipmentAddress, expectedOrders[0].GrandTotal, expectedOrders[0].OrderStatus, expectedOrders[0].PaymentStatus).
		AddRow(expectedOrders[1].ID, expectedOrders[1].CreatedAt, expectedOrders[1].OrderNo, expectedOrders[1].CustomerID, expectedOrders[1].CustomerName, expectedOrders[1].CustomerEmail, expectedOrders[1].ShipmentAddress, expectedOrders[1].GrandTotal, expectedOrders[1].OrderStatus, expectedOrders[1].PaymentStatus)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE customer_id = ? and deleted_at is null ORDER BY id desc")).
		WithArgs(customerID).
		WillReturnRows(rows)

	orders := orderService.GetListByCustomer(customerID)

	if len(orders) != len(expectedOrders) {
		t.Errorf("expected %d orders, got %d", len(expectedOrders), len(orders))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCountOrder(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	orderService := InitOrderService(db)

	totalOrders := int64(5)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `orders`")).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(totalOrders))

	count := orderService.CountOrder()

	if count != totalOrders {
		t.Errorf("expected %d orders, got %d", totalOrders, count)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
