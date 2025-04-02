package dbase

import "database/sql"

func DbConnection() (dbConn *sql.DB) {
	dbConn, err := sql.Open("sqlite3", "./database/beer.db")
	if err != nil {
		dbConn.Close()
		panic(err)
	}
	return dbConn
}
