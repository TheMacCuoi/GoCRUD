package router
import (
	"github.com/labstack/echo"
	
	"goCRUD/service"
)

func Api() {
	e := echo.New()
 	dber := service.UserHandler{}
	//routes
	e.GET("/users", dber.GetAllUsers)
	e.POST("/users", dber.CreateUser)
	e.GET("/users/:id", dber.GetUser)
	e.PUT("/users/:id", dber.UpdateUser)
	e.DELETE("/users/:id", dber.DeleteUser)

	//start server
	e.Start(":1207")
}