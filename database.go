package sccrawler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Initialize(file string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertData(db *sql.DB, r Reservation) error{
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into reservation (id, name, tel, place, date) values (?, ?, ?, ?, ?)")
	_, err := stmt.Exec(r.Id, r.Name, r.Tel, r.Place, r.Date)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}