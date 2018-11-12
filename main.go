package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck/check"
)

var Versions = map[string]func(host string) check.VersionCheck{
	"mysql":      check.NewMysql,
	"postgresql": check.NewPsql,
	"mongo":      check.NewMongo,
	"redis":      check.NewRedis,
}

func NewVersion(db string, host string) check.VersionCheck {
	if _, ok := Versions[db]; ok {
		return Versions[db](host)
	}
	return nil
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	flag.Parse()
	var version = NewVersion(*db, *host)
	fmt.Println(version.Version())
}
