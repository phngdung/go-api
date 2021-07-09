package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/models"
)

const (
	DB_DSN = ""
)

var DB *gorm.DB

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func Connect() {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	fmt.Println("Connected to Postgres!")
	db.AutoMigrate(&models.User{})
	DB = db
	// defer db.Close()
}
