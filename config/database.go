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
	db       *gorm.DB
	err      error
)

func StartDB() {
	host     := os.Getenv("DB_HOST")
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort   := os.Getenv("DB_PORT")
	dbname   := os.Getenv("DB_NAME")

	// db config
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("sucessfully connected to database")
	db.Debug().AutoMigrate(models.Activity{}, models.Scheduler{}, models.Jobs{}, models.Logs{})
}
func GetDB() *gorm.DB {
	return db
}
