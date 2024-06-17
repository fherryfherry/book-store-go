package orderDetails

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

func TestGetListByOrderIds(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	orderDetailService := InitOrderDetailService(db)

	// Define the order IDs to query
	orderIDs := []uint{1, 2}

	// Define the expected order detail models
	expectedOrderDetails := []OrderDetailModel{
		{OrderID: 1, BookCode: "ABC123", BookTitle: "Test Book 1", BookPrice: 19.99, Qty: 2, TotalPrice: 39.98},
		{OrderID: 1, BookCode: "DEF456", BookTitle: "Test Book 2", BookPrice: 29.99, Qty: 1, TotalPrice: 29.99},
		{OrderID: 2, BookCode: "GHI789", BookTitle: "Test Book 3", BookPrice: 9.99, Qty: 3, TotalPrice: 29.97},
	}

	rows := sqlmock.NewRows([]string{"order_id", "book_code", "book_title", "book_price", "qty", "total_price"}).
		AddRow(expectedOrderDetails[0].OrderID, expectedOrderDetails[0].BookCode, expectedOrderDetails[0].BookTitle, expectedOrderDetails[0].BookPrice, expectedOrderDetails[0].Qty, expectedOrderDetails[0].TotalPrice).
		AddRow(expectedOrderDetails[1].OrderID, expectedOrderDetails[1].BookCode, expectedOrderDetails[1].BookTitle, expectedOrderDetails[1].BookPrice, expectedOrderDetails[1].Qty, expectedOrderDetails[1].TotalPrice).
		AddRow(expectedOrderDetails[2].OrderID, expectedOrderDetails[2].BookCode, expectedOrderDetails[2].BookTitle, expectedOrderDetails[2].BookPrice, expectedOrderDetails[2].Qty, expectedOrderDetails[2].TotalPrice)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_details` WHERE order_id IN (?,?)")).
		WithArgs(1, 2).
		WillReturnRows(rows)

	orderDetails := orderDetailService.GetListByOrderIds(orderIDs)

	// Compare the results
	if len(orderDetails) != 2 {
		t.Errorf("expected 2 orders, got %d", len(orderDetails))
	}

	for _, expectedOrderDetail := range expectedOrderDetails {
		orderID := expectedOrderDetail.OrderID
		if len(orderDetails[orderID]) == 0 {
			t.Errorf("expected details for order ID %d, got none", orderID)
		}
		found := false
		for _, actualOrderDetail := range orderDetails[orderID] {
			if actualOrderDetail == expectedOrderDetail {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected order detail %+v not found in result", expectedOrderDetail)
		}
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
