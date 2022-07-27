package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var dbInstance *sql.DB = nil

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	checkErr(err)

	dbInstance = db

	return dbInstance
}

func Instance() *sql.DB {
	return dbInstance
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
