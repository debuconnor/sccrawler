package sccrawler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	checkError(err)

	return db, nil
}

// Query database
// No return data. Use for [insert, delete, update]
func Save(db *sql.DB, query string) {
	_, err := db.Exec(query)
	checkError(err)
}

// Query database
// Result : arr[pk][column]string
func Get(db *sql.DB, query string, key string) map[string]map[string]string {
	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	cols, err := rows.Columns()
	checkError(err)

	result := make(map[string]map[string]string)

	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		data := make(map[string]string)

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		_ = rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
		}

		pk := data[key]
		delete(data, key)

		result[pk] = make(map[string]string)
		result[pk] = data
	}

	return result
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
