package main

import "database/sql"


var db, _ = sql.Open("mysql", "Tung:Tung1272000@tcp(127.0.0.1:3306)/User")