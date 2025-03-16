package main

import (
	"flag"
	"job-scheduler-backend/config"
	"job-scheduler-backend/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Parse command-line flags
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	config.StartDB()

	// Run migrations if the --migrate flag is provided
	if *migrateFlag {
		config.RunMigrations()
	}

	// 🔥 Get DB instance and pass it to handlers
	db := config.GetDB()
	if db == nil {
		panic("Database is not initialized!") // 🚨 Debugging check
	}

	e := echo.New()
	routers.InitRoutes(e)

	port := ":" + os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(port))

}
