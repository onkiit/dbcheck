package main

import (
	"flag"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck/check"
)

var checkers = map[string]func(host string) check.Dialer{
	// "mysql":      check.NewMysql,
	"postgresql": check.NewPsql,
	// "mongo":      check.NewMongo,
	"redis": check.NewRedis,
}

func newChecker(db string, host string) check.Dialer {
	if _, ok := checkers[db]; ok {
		return checkers[db](host)
	}
	return nil
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	flag.Parse()
	var checker = newChecker(*db, *host)
	checker.Dial()
}
