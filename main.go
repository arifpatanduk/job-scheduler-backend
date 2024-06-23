package main

import (
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
	config.StartDB()
	e := echo.New()

	routers.InitRoutes(e)
	port := ":" + os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(port))

}