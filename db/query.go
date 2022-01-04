package db

import (
	"database/sql"
	"goCRUD/ulti"
)

func InsertDB(id int, name string) {
	db := ulti.Init()
	insertDB, err := db.Prepare("INSERT INTO users (id, name) values (?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(id, name)
	defer db.Close()
}
func SelectDB(id int) (result *sql.Rows) {
	db := ulti.Init() 
	var err error
	result, err = db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return result
}
func SelectAllDB() (result *sql.Rows) {
	db := ulti.Init()
	result, _ = db.Query("SELECT * FROM users")
	defer db.Close()
	return result
}
func UpdateDB(id int, name string) {
	db := ulti.Init()
	updateDB, err := db.Prepare("UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateDB.Exec(name, id)
	defer db.Close()
}
func DeleteDB(id int) {
	db := ulti.Init()
	deleteUser, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	deleteUser.Exec(id)
}