package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/globalsign/mgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

type Dialer interface {
	Dial(string) error
}

type VersionCheck interface {
	Version() (string, error)
}

type Mysql struct {
	Host string
}

type Psql struct {
	Host string
}

type Mongo struct {
	Host string
}

type Redis struct {
	Host string
}

func (m Mysql) Version() (string, error) {
	db, err := sql.Open("mysql", m.Host)
	if err != nil {
		fmt.Println("connect mysql", err)
		return "", err
	}
	defer db.Close()

	rows, err := db.Query("SHOW VARIABLES LIKE '%version%'")
	if err != nil {
		fmt.Println("query ", err)
		return "", nil
	}

	var version, variable, value string
	version = "MySql "
	for rows.Next() {
		err := rows.Scan(&variable, &value)
		if err != nil {
			fmt.Println("scan", err)
			return "", err
		}
		version += value + " "
	}

	return version, nil

}

func (p Psql) Version() (string, error) {
	db, err := sql.Open("postgres", p.Host)
	if err != nil {
		return "", err
	}

	defer db.Close()

	var version string
	_ = db.QueryRow("SELECT version()").Scan(&version)

	return version, nil
}

func (m Mongo) Version() (string, error) {
	session, err := mgo.Dial(m.Host)
	if err != nil {
		fmt.Println("mongo conn", err)
		return "", err
	}

	buildInfo, err := session.BuildInfo()
	if err != nil {
		fmt.Println("getting build info", err)
		return "", err
	}

	version := fmt.Sprintf(" MongoDB\n db version %s \n git version %s \n OpenSSL version %s \n", buildInfo.Version, buildInfo.GitVersion, buildInfo.OpenSSLVersion)

	return version, nil
}

func (r Redis) Version() (string, error) {
	con, err := redis.Dial("tcp", r.Host)
	if err != nil {
		return "", err
	}

	version, err := redis.String(con.Do("INFO"))
	if err != nil {
		fmt.Println("getting info", err)
		return "", nil
	}

	return version, nil
}

func DBVersion(db VersionCheck) {
	version, err := db.Version()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(version)
}

func main() {
	db := flag.String("db", "mysql", "Specify your database server")
	host := flag.String("host", "root@tcp(localhost:3306)/test", "Specify your database connection URI depending your server")
	flag.Parse()
	switch *db {
	case "postgresql":
		psq := Psql{
			Host: *host,
		}
		DBVersion(psq)
	case "mongodb":
		mongo := Mongo{
			Host: *host,
		}
		DBVersion(mongo)
	case "redis":
		redis := Redis{
			Host: "localhost:6379",
		}
		DBVersion(redis)
	default:
		msq := Mysql{
			Host: *host,
		}
		DBVersion(msq)
	}

}
