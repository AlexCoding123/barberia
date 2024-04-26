package barbers_routes

import (
	"github.com/labstack/echo/v4"
    "barberia/server/controllers"
)

func RegisterBarberRoutes (e *echo.Echo) {
    e.GET("/barbers", controller.GetBarbers)
    e.GET("/barbers/:id", controller.GetBarberById)
    e.POST("/new-barber", controller.CreateBarber)
    e.PUT("/barbers/:id", controller.EditBarber)

    //Loading Views
    e.GET("/barber-form", controller.GetBarberForm)
    e.GET("/", controller.LoadHomePage)
}
