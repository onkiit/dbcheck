package pgsql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var con *sql.DB

func Open(host string) error {
	db, err := sql.Open("postgree", host)
	if err != nil {
		log.Println("open", err)
	}

	con = db
	return nil
}

func Close() {
	con = nil
}

func DB() *sql.DB {
	return con
}
