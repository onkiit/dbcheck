package sqlite

import (
	"database/sql"
	"fmt"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type sqlite struct {
	db *sql.DB
}

func (s *sqlite) Version() error {
	var version string
	err := s.db.QueryRow("select sqlite_version() as version;").Scan(&version)
	if err != nil {
		return err
	}
	fmt.Printf("SQLite version %s\n", version)
	return nil
}

func (s *sqlite) ActiveClient() error {
	return nil
}

func (s *sqlite) Health() error {
	return nil
}

func (s *sqlite) Dial(host string) dbcheck.Checker {
	db, err := sql.Open("sqlite3", "")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &sqlite{db: db}
}

func init() {
	registry.Register("sqlite", &sqlite{})
}
