package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type sqlite struct {
	db *sql.DB
}

func (s *sqlite) Version() error {
	return nil
}

func (s *sqlite) ActiveClient() error {
	return nil
}

func (s *sqlite) Health() error {
	return nil
}

func (s *sqlite) Dial(host string) dbcheck.Checker {
	db, err := sql.Open("sqlite3", host)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &sqlite{db: db}
}

func init() {
	registry.Register("sqlite", &sqlite{})
}
