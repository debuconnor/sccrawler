package sccrawler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMysql(username, pw, ip, port, dbname string) *sql.DB {
	if ip == "" {
		ip = "127.0.0.1"
	}

	if port == "" {
		port = "3306"
	}

	connString := username + ":" + pw + "@tcp(" + ip + ":" + port + ")/" + dbname
	db, err := sql.Open("mysql", connString)

	checkError(err, "Mysql failed to open database")

	return db
}

func SelectMysql(db *sql.DB, query string, key string) map[string]map[string]string {
	rows, err := db.Query(query)
	checkError(err, "Mysql failed to get rows: "+query)
	defer rows.Close()

	cols, err := rows.Columns()
	checkError(err, "Mysql failed to get columns")

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

func SaveMysql(db *sql.DB, query string) {
	_, err := db.Exec(query)
	checkError(err, "Mysql query failed: "+query)
}

func SelectMysqlTest(db *sql.DB, query string, key string) map[string]map[string]string {
	rows, err := db.Query(query)
	checkError(err, "Mysql failed to get rows: "+query)
	defer rows.Close()

	cols, err := rows.Columns()
	checkError(err, "Mysql failed to get columns")

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
