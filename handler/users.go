package handler

import (
	"fmt"
	"goCRUD/model"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Create User
func (s *Server) CreateUser(c echo.Context) (err error) {
	u := new(model.User)
	s.db.Create(&u)
	return c.JSONPretty(http.StatusCreated, u, " ")
}

//Show single User
func (s *Server) GetUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := getUserOr404(s.db, id)
	if u == nil{
		fmt.Print("User doesn't exist")
		return
	}
	return c.JSONPretty(http.StatusOK, u, " ")
}

//Update User
func (s *Server) UpdateUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := getUserOr404(s.db, id)
	if u == nil{
		fmt.Print("User doesn't exist")
		return
	}
	if err = c.Bind(u); err != nil{
		return
	}
	if err = s.db.Model(&u).Updates(&u).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

//Del User
func (s *Server) DeleteUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := getUserOr404(s.db, id)
	if u == nil{
		fmt.Print("User doesn't exist")
		return
	}
	if err = c.Bind(u); err != nil{
		return
	}
	if err = s.db.Delete(&u, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}

//Show all User
func (s *Server) GetAllUsers(c echo.Context) (err error) {
	u := []model.User{}
	s.db.Find(&u)
	return c.JSONPretty(http.StatusOK, &u, " ")
}

// select
func getUserOr404(db *gorm.DB, id int) (*model.User, *echo.HTTPError) {
	s := &model.User{}
	if err := db.First(&s, id).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return s, nil
}
