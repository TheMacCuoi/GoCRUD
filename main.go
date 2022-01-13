package main

import (
	"goCRUD/db"
	"goCRUD/router"
)

func main() {
	db.Init()
	router.Api()
}