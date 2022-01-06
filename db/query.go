package db

import (
	"database/sql"
	"goCRUD/ulti"
)

type SQLRepo struct {
	
}

func Init() (db *sql.DB) {
	config, _ := ulti.LoadConfig(".")
	db, _ = sql.Open(config.DBDriver, config.DBSource)
	return db
}

func (repo *SQLRepo) InsertDB(id int, name string) {
	db := Init()
	insertDB, err := db.Prepare("INSERT INTO Users (id, name) values (?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(id, name)
	defer db.Close()
}
func (repo *SQLRepo) SelectDB(id int) (result *sql.Rows) {
	db := Init() 
	var err error
	result, err = db.Query("SELECT * FROM Users WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return result
}
func (repo *SQLRepo) SelectAllDB() (result *sql.Rows){
	db := Init()
	result, _ = db.Query("SELECT * FROM Users")
	defer db.Close()
	return result
}
func (repo *SQLRepo) UpdateDB(id int, name string) {
	db := Init()
	updateDB, err := db.Prepare("UPDATE Users SET name = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateDB.Exec(name, id)
	defer db.Close()
}
func (repo *SQLRepo) DeleteDB(id int) {
	db := Init()
	deleteUser, err := db.Prepare("DELETE FROM Users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	deleteUser.Exec(id)
}