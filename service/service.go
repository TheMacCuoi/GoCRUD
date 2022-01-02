package service

import (
	"database/sql"
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
	db, _ = sql.Open("mysql", "Tung:Tung1272000@tcp(127.0.0.1:3306)/User")
)

//Create User
func CreateUser(c echo.Context) error {
	u := &user{
		ID: idm,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	insertDB, err := db.Prepare("INSERT INTO users (id, name) values (?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(u.ID, u.Name)
	idm++
	defer insertDB.Close()
	return c.JSON(http.StatusCreated, u)
}

//Show single User
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, _ := db.Query("SELECT * FROM users WHERE id = ?", id)
	var u user
	_ = result.Scan(&u.ID, &u.Name)
	defer result.Close()
	return c.JSON(http.StatusOK, u)
}

//Update User
func UpdateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	updateDB, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateDB.Exec(u.Name, id)
	defer updateDB.Close()
	return c.JSON(http.StatusOK, u)
}

//Del User
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	deleteUser, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	deleteUser.Exec(id)
	u := "deleted user id: " + strconv.Itoa(id)
	return c.JSON(http.StatusOK, u)
}

//Show all user
func GetAllUsers(c echo.Context) error {
	var sliceUsers []user
	result, _ := db.Query("SELECT * FROM users")
	for result.Next() {
		var u user
		_ = result.Scan(&u.ID, &u.Name)
		sliceUsers = append(sliceUsers, u)
	}
	return c.JSON(http.StatusOK, sliceUsers)
}
