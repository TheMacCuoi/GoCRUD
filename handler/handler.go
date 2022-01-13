package handler

import (
	"goCRUD/db"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
var (
	idm   = 5
)
type handler interface{
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
}
type UserHandler struct {
	dber db.UserRepo
}
//Create User
func (d *UserHandler) CreateUser(c echo.Context) error {
	u,_ := d.dber.InsertUser(c)
	return c.JSON(http.StatusCreated, u)
}

//Show single User
func (d *UserHandler) GetUser(c echo.Context) error {
	//var sliceUsers []User
	result, _ := d.dber.SelectUser(c)
	/*for result.Next() {
		var u User
		_ = result.Scan(&u.ID, &u.Name)
		sliceUsers = append(sliceUsers, u)
	}*/
	return c.JSON(http.StatusOK, result)
}

//Update User
func (d *UserHandler) UpdateUser(c echo.Context) error {
	u, _ := d.dber.UpdateUser(c)
	return c.JSON(http.StatusOK, u)
}

//Del User
func (d *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	d.dber.DeleteUser(c)
	u := "deleted User id: " + strconv.Itoa(id)
	return c.JSON(http.StatusOK, u)
}

//Show all User
func (d *UserHandler) GetAllUsers(c echo.Context) error {
	result,_ := d.dber.SelectAllUser(c)
	return c.JSON(http.StatusOK, result)
}
