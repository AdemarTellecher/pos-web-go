package dbase

import (
	"database/sql"
	"log"
)

func DbConnection() (dbConn *sql.DB, err error) {
	dbConn, err = sql.Open("sqlite3", "./database/beer.db")
	if err != nil {
		dbConn.Close()
		log.Fatal(err.Error())
	}
	return dbConn, err
}
