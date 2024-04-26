package controller

import (
	"net/http"
	"barberia/server/config"
	"barberia/server/models"

	"github.com/labstack/echo/v4"
)


func CreateBarber(c echo.Context) error{
   	b := new(models.Barber)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

        return c.JSON(http.StatusInternalServerError, data)
    }

    barber := &models.Barber{
        Name:        b.Name,
        LastName:    b.LastName,
        BarberName:  b.BarberName,
        PhoneNumber: b.PhoneNumber,
        Password:    b.Password,
    }

    if err := db.Create(&barber).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusOK, response)
}

func GetBarberForm(c echo.Context) error{
    return c.Render(200, "new-barber", nil)
}


func LoadHomePage(c echo.Context) error{
    return c.Render(200, "index", nil)
}

func GetBarbers(c echo.Context) error{
	db := config.DB()
    var barbers []models.Barber

    if err := db.Find(&barbers).Error; err != nil{
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
    }

    response := map[string]interface{}{
		"data": barbers,
	}

	return c.JSON(http.StatusOK,response)
}

func GetBarberById(c echo.Context) error{
	db := config.DB()
    var barber models.Barber
    id := c.Param("id")

    if err := db.Find(&barber, id).Error; err != nil{
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
    }

    response := map[string]interface{}{
		"data": barber,
	}

	return c.JSON(http.StatusOK, response)
}

func EditBarber(c echo.Context) error {
	db := config.DB()
    var barber models.Barber
    id := c.Param("id")

    if err := db.Find(&barber, id).Error; err != nil{
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
    }

	// Binding data
	if err := c.Bind(&barber); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

        return c.JSON(http.StatusInternalServerError, data)
    }

    if err := db.Save(&barber).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": barber,
	}

	return c.JSON(http.StatusOK, response)
}
