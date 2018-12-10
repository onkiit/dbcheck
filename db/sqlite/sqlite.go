package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/onkiit/dbinfo"
	sq "github.com/onkiit/dbinfo/sqlite"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type sqlite struct {
	db *sql.DB
}

func (s *sqlite) Version() error {
	con := &dbinfo.Conn{
		DB: s.db,
	}
	store := sq.New(con)
	v, err := store.GetVersion(context.Background())
	if err != nil {
		return err
	}

	fmt.Print(v.Version)
	return nil
}

func (s *sqlite) ActiveClient() error {
	fmt.Println("Sqlite does not provided active_client checking")
	return nil
}

func (s *sqlite) Health() error {
	con := &dbinfo.Conn{
		DB: s.db,
	}

	store := sq.New(con)
	h, err := store.GetHealth(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("health_status: \n pragma_page_size: %d\n pragma_page_count: %d\n", h.SQLiteHealth.PragmaPageSize, h.SQLiteHealth.PragmaPageCount)

	return nil
}

func (s *sqlite) Dial(path string) dbcheck.Checker {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &sqlite{db: db}
}

func init() {
	registry.Register("sqlite", &sqlite{})
}
