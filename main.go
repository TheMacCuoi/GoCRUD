package main

import (
	"goCRUD/db"
	"goCRUD/handler"
)

func main() {
	db.Init()
	e := handler.Api()

	//start server
	e.Start(":1207")
}
