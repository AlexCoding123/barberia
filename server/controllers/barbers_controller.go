package controller

import (
	"net/http"
	"barberia/server/config"
	"barberia/server/models"
    "fmt"
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
        ProfilePicture: b.ProfilePicture,
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
    fmt.Println("LoadHomePage function called")
    // Fetch barbers using GetBarbers
    barbers, err := GetBarbers(c)
    if err != nil {
        // Handle error if GetBarbers fails
        fmt.Println(err)
        return err
    }
    return c.Render(http.StatusOK, "index", barbers)
}

func GetBarbers(c echo.Context) ([]models.Barber, error){
	db := config.DB()
    var barbers []models.Barber

    if err := db.Find(&barbers).Error; err != nil{
		return nil, err
    }

	return barbers, nil
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
