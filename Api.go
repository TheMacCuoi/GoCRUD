package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//routes
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	//start server
	e.Start(":1207")
}