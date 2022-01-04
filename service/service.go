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

//Create User
func CreateUser(c echo.Context) error {
	u := &user{
		ID: idm,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	db.InsertDB(u.ID, u.Name)
	idm++
	return c.JSON(http.StatusCreated, u)
}

//Show single User
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var sliceUsers []user
	result := db.SelectDB(id)
	for result.Next() {
		var u user
		_ = result.Scan(&u.ID, &u.Name)
		sliceUsers = append(sliceUsers, u)
	}
	return c.JSON(http.StatusOK, sliceUsers)
}

//Update User
func UpdateUser(c echo.Context) error {
	
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	db.UpdateDB(id, u.Name)
	return c.JSON(http.StatusOK, u)
}

//Del User
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db.DeleteDB(id)
	u := "deleted user id: " + strconv.Itoa(id)
	return c.JSON(http.StatusOK, u)
}

//Show all user
func GetAllUsers(c echo.Context) error {
	var sliceUsers []user
	result := db.SelectAllDB()
	for result.Next() {
		var u user
		_ = result.Scan(&u.ID, &u.Name)
		sliceUsers = append(sliceUsers, u)
	}
	return c.JSON(http.StatusOK, sliceUsers)
}
