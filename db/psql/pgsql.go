package psql

import (
	"database/sql"
	"errors"
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

	info := "health_status: \n Database Information \n"
	for rows.Next() {
		if err := rows.Scan(&datname, &size); err != nil {
			return err
		}
		info += "  DB Name: " + datname + "\t\tSize: " + size + "\n"
	}
	fmt.Print(info)

	if err := p.getTableSize(); err != nil {
		return err
	}
	return nil
}

func (p *psql) getTables() (map[string][]string, error) {
	rows, err := p.DB.Query("select schemaname as schema, relname as table from pg_statio_all_tables where schemaname not in ('pg_catalog', 'pg_toast', 'information_schema')")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tables := make(map[string][]string)
	for rows.Next() {
		var schema, table string
		err := rows.Scan(&schema, &table)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tables[schema] = append(tables[schema], table)
	}

	if len(tables) < 1 {
		return nil, errors.New("Could not find table")
	}

	return tables, nil
}

func (p *psql) getTableSize() error {
	tables, err := p.getTables()
	if err != nil {
		return err
	}
	fmt.Println(" Table Information")
	for k, v := range tables {
		var tableSize, indexSize string
		fmt.Printf("  > Schema: %s\n", k)
		if len(v) > 0 {
			for _, val := range v {
				qTable := fmt.Sprintf("SELECT pg_size_pretty(pg_total_relation_size('%s.%s')) as tableSize", k, val)
				qIndex := fmt.Sprintf("SELECT pg_size_pretty(pg_indexes_size('%s.%s')) as indexSize", k, val)
				err := p.DB.QueryRow(qTable).Scan(&tableSize)
				if err != nil {
					fmt.Println(err)
					return err
				}
				err = p.DB.QueryRow(qIndex).Scan(&indexSize)
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Printf("     Table: %s\n      Table Size: %s\n      Index Size: %s\n", val, tableSize, indexSize)
			}
		}
	}
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
	registry.Register("postgresql", &psql{})
}
