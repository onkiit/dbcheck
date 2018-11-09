package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Dialer interface {
	Dial() error
	Close() error
}

type VersionCheck interface {
	Version() (string, error)
}

type Mysql struct {
	Dbserver string
	Host     string
}

type Psql struct {
	Host string
}

func (m Mysql) Version() (string, error) {
	out, err := exec.Command(m.Dbserver, "-V").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (p Psql) Version() (string, error) {
	db, err := sql.Open("postgres", p.Host)
	if err != nil {
		return "", err
	}

	err = db.Ping()
	stmt, err := db.Prepare("SELECT version()")
	if err != nil {
		return "", err
	}

	var version string
	_ = stmt.QueryRow().Scan(&version)

	return version, nil
}

func DBVersion(db VersionCheck) {
	version, err := db.Version()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(version)
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root:root@localhost/user-db", "Specify your database connection URI depending your server")
	flag.Parse()
	switch *db {
	case "postgresql":
		psq := Psql{
			Host: *host,
		}
		DBVersion(psq)
	default:
		msq := Mysql{
			Dbserver: *db,
			Host:     *host,
		}
		DBVersion(msq)
	}

}
