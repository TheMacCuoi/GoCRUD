package db

import (
	"fmt"
	"goCRUD/model"
	"goCRUD/ulti"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Init(){
	configuration, _ := ulti.LoadConfig(".")
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", configuration.DB_USER, configuration.DB_PASS, configuration.DB_NAME)
	db, err := gorm.Open(mysql.Open(connect_string), &gorm.Config{})
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.User{})
}

func GetConn() (db *gorm.DB, err error){
	configuration, _ := ulti.LoadConfig(".")
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", configuration.DB_USER, configuration.DB_PASS, configuration.DB_NAME)
	db, err = gorm.Open(mysql.Open(connect_string), &gorm.Config{})
	return db, nil
}