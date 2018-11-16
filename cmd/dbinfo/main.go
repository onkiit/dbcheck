package main

import (
	c "github.com/onkiit/dbcheck/checkdb"
)

/*
REVIEW
1. naming package
2. implement registry pattern
3. handling possible nil return function from the caller
4. folder structure: https://github.com/golang-standards/project-layout
5.
*/

type App struct {
	DBCheck map[string]c.Dialer
}

func main() {
	app := App{
		DBCheck: make(map[string]c.Dialer),
	}
	post := app.DBCheck["postgresql"]

	a := post.Dial("postgresql://postgres:postgres@localhost/postgres")
	if err := a.Version(); err != nil {
		return
	}
	if err := a.ActiveClient(); err != nil {
		return
	}
	if err := a.Health(); err != nil {
		return
	}

	// db := flag.String("db", "mysql", "Specify your database server")
	// host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	// flag.Parse()
	// var checker = newChecker(*db, *host)
	// if checker == nil {
	// 	fmt.Printf("(%s) Unsupported database.\n", *db)
	// 	return
	// }
	// checker.GetInfo()
}
