package main

import (
	"fmt"
	"goCRUD/db"
	"goCRUD/handler"
	"goCRUD/model"
)

func main() {
	db, err := db.GetConn()
	if err != nil {
		fmt.Print("Failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	handler.Api(db)

}
