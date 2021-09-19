package routes

import (
	"golang-crud/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/user", controllers.GetUsersAllController)
	e.GET("/user/:id", controllers.GetUsersByIDController)
	e.POST("/user", controllers.CreateUsersController)
	e.PUT("/user/:id", controllers.UpdateUsersController)
	e.DELETE("/user/:id", controllers.DeleteUsersController)

	return e
}
