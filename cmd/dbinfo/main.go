package main

import (
	"flag"
	"fmt"

	_ "github.com/onkiit/dbcheck/db/bolt"
	_ "github.com/onkiit/dbcheck/db/cassandra"
	_ "github.com/onkiit/dbcheck/db/mongo"
	_ "github.com/onkiit/dbcheck/db/mysql"
	_ "github.com/onkiit/dbcheck/db/psql"
	_ "github.com/onkiit/dbcheck/db/redis"
	_ "github.com/onkiit/dbcheck/db/sqlite"
	"github.com/onkiit/dbcheck/registry"
)

func dbInfo(db string, host string, path string) {
	var dial string
	if host == "" {
		dial = path
	} else {
		dial = host
	}
	dialer := registry.Dialers(db)
	if dialer == nil {
		fmt.Printf("(%s) Database not supported\n", db)
		return
	}
	checker := dialer.Dial(dial)
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
	db := flag.String("db", "", "Specify your database server. Supported databases (key): redis, mongo, postgresql, mysql, cassandra, bolt, sqlite ")
	host := flag.String("host", "", "Specify your database connection URI depending your server")
	path := flag.String("path", "", "Specify your database path (used for bolt and sqlite)")
	flag.Parse()

	dbInfo(*db, *host, *path)
}
