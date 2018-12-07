package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/onkiit/dbinfo"

	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck"
	"github.com/onkiit/dbcheck/registry"
	"github.com/onkiit/dbinfo/postgres"
)

type psql struct {
	DB *sql.DB
}

func (p *psql) Version() error {
	con := &dbinfo.Conn{
		DB: p.DB,
	}
	store := postgres.New(con)
	res, err := store.GetVersion(context.Background())
	if err != nil {
		return err
	}

	fmt.Println(res.Version)
	return nil
}

func (p *psql) ActiveClient() error {
	con := &dbinfo.Conn{
		DB: p.DB,
	}
	store := postgres.New(con)
	res, err := store.GetActiveClient(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("active_client(s): %d\n", res.ActiveClient)
	return nil
}

func (p *psql) Health() error {
	con := &dbinfo.Conn{
		DB: p.DB,
	}
	store := postgres.New(con)
	res, err := store.GetHealth(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("health_status: ")
	fmt.Printf(" Database Information: \n  Database Name: %s\t Database Size: %s\n", res.PsqlHealth.DBInformation.DBName, res.PsqlHealth.DBInformation.DBSize)
	fmt.Println(" Table Information: ")
	for _, v := range res.PsqlHealth.TableInformation {
		fmt.Printf("  > Schema: %s\n    Table: %s\n    Table Size: %s\n    Index Size: %s\n", v.SchemaName, v.TableName, v.TableSize, v.IndexSize)
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
