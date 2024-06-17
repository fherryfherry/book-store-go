package main

import (
	"booking-online/commons/jwt"
	validatorCommon "booking-online/commons/validator"
	"booking-online/handlers/book"
	"booking-online/handlers/login"
	"booking-online/handlers/order"
	"booking-online/handlers/registration"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file")
	}
}

func initDb() *gorm.DB {
	username := viper.GetString("database.user")
	password := viper.GetString("database.password")
	hostname := viper.GetString("database.host")
	dbName := viper.GetString("database.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	initConfig()

	e := echo.New()

	// Enable validator json field
	e.Validator = &validatorCommon.CustomValidator{Validator: validator.New()}

	// Enable logging
	e.Use(middleware.Logger())

	// Connection for Databases
	db := initDb()

	// Init handler
	registerHandler := registration.InitHandler(db)
	loginHandler := login.InitHandler(db)
	bookHandler := book.InitHandler(db)
	orderHandler := order.InitHandler(db)

	// Routing for non secured area
	e.POST("/v1/registration", registerHandler.RegisterHandler)
	e.POST("/v1/login", loginHandler.LoginHandler)

	// Routing for secured Area
	r := e.Group("/secured/v1")
	r.Use(jwt.InitMiddlewareJwt())
	r.GET("/book/list", bookHandler.GetBookListHandler)
	r.POST("/book/order", orderHandler.CheckoutOrder)
	r.GET("/order/my-history", orderHandler.MyOrder)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))))
}
