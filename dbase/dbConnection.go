package dbase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DbConnection() (dbConn *sql.DB, err error) {
	dbConn, err = sql.Open("sqlite3", "./database/beer.db")
	return dbConn, err
}
