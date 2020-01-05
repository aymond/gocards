package main

import (
	"database/sql"
	"log"

	// Need to pass a driver name as the first argument of the sql.Open() function
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "dbname=game.db")
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()
	return
}

func dbCreate() {
	var err error
	Db, err = sql.Open("sqlite3", "dbname=game.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt, _ := Db.Prepare("CREATE TABLE IF NOT EXISTS games (id INTEGER SERIAL PRIMARY KEY, uuid TEXT, game_name TEXT, created_at TIMESTAMP NOT NULL)")
	stmt.Exec()

	defer Db.Close()
	return
}
