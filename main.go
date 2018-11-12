package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/onkiit/dbcheck/database"
)

func DBVersion(db database.VersionCheck) {
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
	var versionCheck database.VersionCheck
	switch *db {
	case "postgresql":
		versionCheck = database.NewPsql(*host)
	case "mongodb":
		versionCheck = database.NewMongo(*host)
	case "redis":
		versionCheck = database.NewRedis(*host)
	default:
		versionCheck = database.NewMysql(*host)
	}

	DBVersion(versionCheck)

}
