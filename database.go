package sccrawler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Save(file string, r Reservation) error{
	db, err := sql.Open("sqlite3", file)
	stmt, _ := db.Prepare("insert into reservation (id, name, tel, place, date) values (?, ?, ?, ?, ?)")
	stmt.Exec(r.Id, r.Name, r.Tel, r.Place, r.Date)
	
	return err
}