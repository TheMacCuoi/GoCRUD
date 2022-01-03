package router
import (
	"github.com/labstack/echo"
	
	service "goCRUD/service"
)

func Api() {
	e := echo.New()

	//routes
	e.GET("/users", service.GetAllUsers)
	e.POST("/users", service.CreateUser)
	e.GET("/users/:id", service.GetUser)
	e.PUT("/users/:id", service.UpdateUser)
	e.DELETE("/users/:id", service.DeleteUser)

	//start server
	e.Start(":1207")
}