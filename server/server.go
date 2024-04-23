package main

import (
	"barberia/server/config"
	"barberia/server/routes"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    barbers_routes.RegisterBarberRoutes(e)
    // Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	e.Logger.Fatal(e.Start(":8080"))
}
