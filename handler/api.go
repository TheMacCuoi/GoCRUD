package handler

import (
	"goCRUD/repository"

	"github.com/labstack/echo"
)

type UserHandler struct {
	dber repository.UserRepo
}

func Api() *echo.Echo {
	e := echo.New()
	dber := UserHandler{}
	//routes
	e.GET("/users", dber.GetAllUsers)
	e.POST("/users", dber.CreateUser)
	e.GET("/users/:id", dber.GetUser)
	e.PUT("/users/:id", dber.UpdateUser)
	e.DELETE("/users/:id", dber.DeleteUser)

	return e
}
