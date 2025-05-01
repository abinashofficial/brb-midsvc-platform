package config

import (
	"brb-midsvc-platform/model"
	"fmt"
	"log"
	"os"
    "github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func SetupDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)      // max concurrent connections
	sqlDB.SetMaxIdleConns(10)       // idle connections
	sqlDB.SetConnMaxLifetime(time.Hour)


  // Correct order: Vendor first, then Service
  err = db.AutoMigrate(
    &model.Vendor{},
    &model.Service{},
    &model.Booking{},
	&model.LinkServiceVendor{},
)
if err != nil {
    log.Fatal("Failed auto migration:", err)
}

	return db
}

func HealthCheck(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(500, gin.H{"status": "Database Connection Error"})
			return
		}
		c.JSON(200, gin.H{"status": "OK"})
	}
}
