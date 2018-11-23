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
	if dialer == nil {
		fmt.Printf("(%s) Database not supported\n", db)
		return
	}
	checker := dialer.Dial(host)
	if checker == nil {
		fmt.Println("Server unreachable")
		return
	}
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
	db := flag.String("db", "redis", "Specify your database server. Supported databases (key): redis, mongo, postgresql, mysql ")
	host := flag.String("host", "localhost:6379", "Specify your database connection URI depending your server")
	flag.Parse()

	dbInfo(*db, *host)
}
