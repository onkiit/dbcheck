package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
)

type psql struct {
	DB *sql.DB
}

func (p *psql) Version() error {
	var version string
	_ = p.DB.QueryRow("SELECT version()").Scan(&version)

	fmt.Println(version)
	return nil
}

func (p *psql) ActiveClient() error {
	var count int
	err := p.DB.QueryRow("SELECT count(0) FROM pg_stat_activity where state='active' ").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return err
	}

	info := fmt.Sprintf("active_client(s): %d", count)
	fmt.Println(info)
	return nil
}

func (p *psql) Health() error {
	var datname, size string
	rows, err := p.DB.Query("select datname, pg_size_pretty(pg_database_size(datname)) as size from pg_database order by pg_database_size(datname) desc;")
	if err != nil {
		return err
	}

	info := "health_status: \n Storage Information \n"
	for rows.Next() {
		if err := rows.Scan(&datname, &size); err != nil {
			return err
		}
		info += " DB Name: " + datname + "     Size: " + size + "\n"
	}

	fmt.Println(info)
	return nil
}

func (p *psql) Dial(host string) dbcheck.Checker {
	db, err := sql.Open("postgres", host)
	if err != nil {
		return nil
	}

	return &psql{DB: db}
}

func init() {
	registry.RegisterDialers("postgresql", &psql{})
}
