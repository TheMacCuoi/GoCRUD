package handler

import (

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}


func Api(db *gorm.DB) {
	e := echo.New()
	s := &Server{db: db}
	//routes
	e.GET("/users", s.GetAllUsers)
	e.POST("/users", s.CreateUser)
	e.GET("/users/:id", s.GetUser)
	e.PUT("/users/:id", s.UpdateUser)
	e.DELETE("/users/:id", s.DeleteUser)

	//start server
	e.Start(":1207")
}
