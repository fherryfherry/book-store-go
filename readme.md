# Book Store Go
by Ferry Ariawan

### Tech Stack:
- Language: Golang 1.22
- DB: MySQL

### Dependencies:
- github.com/go-playground/validator v9.31.0+incompatible
- github.com/golang-jwt/jwt/v5 v5.0.0
- github.com/labstack/echo-jwt/v4 v4.2.0
- github.com/labstack/echo/v4 v4.12.0
- github.com/spf13/viper v1.19.0
- github.com/stretchr/testify v1.9.0
- golang.org/x/crypto v0.24.0
- gorm.io/driver/mysql v1.5.7
- gorm.io/gorm v1.25.10

### Installation
1. Import the database named `book_store.sql` into your mySQL local server
1. Config the database credential on `config.yaml` at the root directory
1. Update dependencies by run command `go mod tidy`

### Run Project
```bash
go run main.go
```

### Run Unit Test
```bash
go test -v ./...
```

### Postman Collection
Import postman file named `BookStore.postman_collection.json`