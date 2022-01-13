package db

import (
	"goCRUD/model"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)
var u model.User

type UserRepo struct {
	db *gorm.DB
}
func New(db *gorm.DB) UserRepo{
	return UserRepo{
		db: db,
	}
}
func (repo *UserRepo) InsertUser(c echo.Context) (*gorm.DB, error){
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return repo.db.Create(&u), nil
}
func (repo *UserRepo) SelectUser(c echo.Context) (*gorm.DB, error){
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return repo.db.First(&u, id), nil
}
func (repo *UserRepo) SelectAllUser(c echo.Context) (*gorm.DB, error){
	return repo.db.Select(&u), nil
}
func (repo *UserRepo) UpdateUser(c echo.Context) (*gorm.DB, error){
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return repo.db.Model(&u).Updates(map[string]interface{}{
		"id": u.ID,
		"name": u.Name,
	}), nil
}
func (repo *UserRepo) DeleteUser(c echo.Context) (*gorm.DB, error){
	if err := c.Bind(&u); err != nil {
		return nil, err
	}
	return repo.db.Delete(&u, u.ID), nil
}
