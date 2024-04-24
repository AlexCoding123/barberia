package barbers_routes

import (
	"github.com/labstack/echo/v4"
    "barberia/server/controllers"
)

func RegisterBarberRoutes (e *echo.Echo) {
    e.GET("/barbers", controller.GetBarbers)
    e.GET("/barbers/:id", controller.GetBarberById)
    e.POST("/barbers", controller.CreateBarber)
}
