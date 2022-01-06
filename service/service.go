package service

import (
	"goCRUD/db"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
var (
	idm   = 5
)
type UserHandler struct {
	dber db.SQLRepo
}
//Create User
func (d *UserHandler) CreateUser(c echo.Context) error {
	u := &user{
		ID: idm,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	d.dber.InsertDB(u.ID, u.Name)
	idm++
	return c.JSON(http.StatusCreated, u)
}

//Show single User
func (d *UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var sliceUsers []user
	result := d.dber.SelectDB(id)
	for result.Next() {
		var u user
		_ = result.Scan(&u.ID, &u.Name)
		sliceUsers = append(sliceUsers, u)
	}
	return c.JSON(http.StatusOK, sliceUsers)
}

//Update User
func (d *UserHandler) UpdateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	d.dber.UpdateDB(id, u.Name)
	return c.JSON(http.StatusOK, u)
}

//Del User
func (d *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	d.dber.DeleteDB(id)
	u := "deleted user id: " + strconv.Itoa(id)
	return c.JSON(http.StatusOK, u)
}

//Show all user
func (d *UserHandler) GetAllUsers(c echo.Context) error {
	var sliceUsers []user
	result := d.dber.SelectAllDB()
	for result.Next() {
		var u user
		_ = result.Scan(&u.ID, &u.Name)
		sliceUsers = append(sliceUsers, u)
	}
	return c.JSON(http.StatusOK, sliceUsers)
}
