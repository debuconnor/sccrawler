package sccrawler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(file string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", file)
	checkError(err)

	return db, nil
}

func Save(db *sql.DB, r Reservation) error{
	tx, _ := db.Begin()
	statement, _ := db.Prepare("insert into reservation (id, name, tel, place, date) values (?, ?, ?, ?, ?)")
	_, err := statement.Exec(r.Id, r.Name, r.Tel, r.Place, r.Date)
	checkError(err)

	tx.Commit()
	return nil
}

func Get(db *sql.DB, query string, key string) map[string]map[string]string{	
	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	cols, err := rows.Columns()
	checkError(err)

	result := make(map[string]map[string]string)

	for rows.Next(){
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		data := make(map[string]string)
		
		for i := range columns{
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols{
			data[colName] = columns[i]
		}

		id := data[key]
		delete(data, key)

		result[id] = make(map[string]string)
		result[id] = data
	}

	return result
}

func checkError(e error){
	if e != nil{
		panic(e)
	}
}