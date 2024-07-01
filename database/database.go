package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql" // Import the GORM driver for your database
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func init() {
	Init()
}

func Init() *gorm.DB {
	godotenv.Load()

	if os.Getenv("DB_TYPE") == "" {
		log.Fatal("DB_TYPE is not set")
	} else if os.Getenv("DB_TYPE") == "mysql" {
		if os.Getenv("DB_USER") == "" {
			log.Fatal("DB_USER is not set")
		}
		if os.Getenv("DB_PASSWORD") == "" {
			log.Fatal("DB_PASSWORD is not set")
		}
		if os.Getenv("DB_HOST") == "" {
			log.Fatal("DB_HOST is not set")
		}
		if os.Getenv("DB_PORT") == "" {
			log.Fatal("DB_PORT is not set")
		}
		if os.Getenv("DB_NAME") == "" {
			log.Fatal("DB_NAME is not set")
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		DB = db
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Failed to access underlying DB connection: %v", err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		
	} else if os.Getenv("DB_TYPE") == "sqlite3" {
		if os.Getenv("DB_PATH") == "" {
			log.Fatal("DB_PATH is not set")
		}
		db, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		DB = db
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Failed to access underlying DB connection: %v", err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}


	return DB
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to access underlying DB connection: %v", err)
	}
	sqlDB.Close()
}
