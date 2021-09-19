package controllers

import (
	"golang-crud/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersAllController(c echo.Context) error {
	users, err := database.GetUsersAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "err",
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   newResponseArray(*users),
	})
}

func GetUsersByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := database.GetUsersByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "err",
			"message": "not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   newResponse(*users),
	})
}

func CreateUsersController(c echo.Context) error {
	var userReq UserRequest
	c.Bind(&userReq)

	users, err := database.CreateUsers(userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "err",
			"message": "bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   newResponse(*users),
	})
}

func UpdateUsersController(c echo.Context) error {
	var userReq UserRequest
	c.Bind(&userReq)
	id, _ := strconv.Atoi(c.Param("id"))

	users, err := database.UpdateUsers(id, userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "err",
			"message": "not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "data was updated",
		"data":   newResponse(*users),
	})
}

func DeleteUsersController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, find_err := database.GetUsersByID(id)
	_, err := database.DeleteUsers(id, *users)

	if find_err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "err",
			"message": "id not found",
		})
	} else if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "err",
			"message": "bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
		//"data":   newResponse(*result),
	})
}
