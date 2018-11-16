package main

import (
	"flag"
	"fmt"

	_ "github.com/onkiit/dbcheck/db/mongo"
	_ "github.com/onkiit/dbcheck/db/mysql"
	_ "github.com/onkiit/dbcheck/db/psql"
	_ "github.com/onkiit/dbcheck/db/redis"
	"github.com/onkiit/dbcheck/registry"
)

func dbInfo(db string, host string) {
	dialer := registry.Dialers(db)
	checker := dialer.Dial(host)
	if err := checker.Version(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checker.ActiveClient(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checker.Health(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	flag.Parse()
	dbInfo(*db, *host)
}
