package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck/check"
)

var versions = map[string]func(host string) check.VersionChecker{
	"mysql":      check.NewMysql,
	"postgresql": check.NewPsql,
	"mongo":      check.NewMongo,
	"redis":      check.NewRedis,
}

func newVersion(db string, host string) check.VersionChecker {
	if _, ok := versions[db]; ok {
		return versions[db](host)
	}
	return nil
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	flag.Parse()
	var version = newVersion(*db, *host)
	fmt.Println(version.Version())
}
