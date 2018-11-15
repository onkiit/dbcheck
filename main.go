package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck/check"
)

/*
REVIEW
1. naming package
2. implement registry pattern
3. handling possible nil return function from the caller
4. folder structure: https://github.com/golang-standards/project-layout
5.
*/

var checkers = map[string]func(host string) check.DBChecker{
	"mysql":      check.NewMysql,
	"postgresql": check.NewPsql,
	"mongo":      check.NewMongo,
	"redis":      check.NewRedis,
}

func newChecker(db string, host string) check.DBChecker {
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
	if checker == nil {
		fmt.Printf("(%s) Unsupported database.\n", *db)
		return
	}
	checker.GetInfo()
}
