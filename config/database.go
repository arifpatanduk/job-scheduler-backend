package config

import (
	"fmt"
	"job-scheduler-backend/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// db config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, dbPort, dbname,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("successfully connected to database")
}

func RunMigrations() {
	if db == nil {
		log.Fatal("database connection is not initialized")
	}
	db.Debug().AutoMigrate(models.Activity{}, models.Scheduler{}, models.Jobs{}, models.Log{})
	fmt.Println("migrations completed successfully")
}

func GetDB() *gorm.DB {
	return db
}
