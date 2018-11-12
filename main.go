package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck/check"
)

func DBVersion(db check.VersionCheck) {
	version, err := db.Version()
	if err != nil {
		fmt.Println("print version", err)
		return
	}
	fmt.Println(version)
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	flag.Parse()
	var versionCheck check.VersionCheck
	switch *db {
	case "postgresql":
		versionCheck = check.NewPsql(*host)
	case "mongodb":
		versionCheck = check.NewMongo(*host)
	case "redis":
		versionCheck = check.NewRedis(*host)
	default:
		versionCheck = check.NewMysql(*host)
	}

	DBVersion(versionCheck)

}
