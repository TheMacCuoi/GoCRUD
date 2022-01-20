package repository

import (
	"goCRUD/db"
	"goCRUD/model"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var u model.User

type UserRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}
func (repo *UserRepo) InsertUser(c echo.Context) (*gorm.DB, error) {
	data, _ := db.GetConn()
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return data.Create(u), nil
}
func (repo *UserRepo) SelectUser(c echo.Context) (*gorm.DB, error) {
	data, _ := db.GetConn()
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return data.First(u, id), nil
}
func (repo *UserRepo) SelectAllUser(c echo.Context) (*gorm.DB, error) {
	data, _ := db.GetConn()
	return data.Select(u), nil
}
func (repo *UserRepo) UpdateUser(c echo.Context) (*gorm.DB, error) {
	data, _ := db.GetConn()
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return data.Model(u).Updates(map[string]interface{}{
		"id":   u.ID,
		"name": u.Name,
	}), nil
}
func (repo *UserRepo) DeleteUser(c echo.Context) (*gorm.DB, error) {
	data, _ := db.GetConn()
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return data.Delete(u, u.ID), nil
}
