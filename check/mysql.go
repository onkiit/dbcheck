package check

import (
	"database/sql"
	"fmt"
)

type mysql struct {
	host string
}

func (m mysql) Version() (string, error) {
	db, err := sql.Open("mysql", m.host)
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

func (m mysql) ActiveClient() (string, error) {
	return "", nil
}

func (m mysql) Health() (string, error) {
	return "", nil
}

func NewMysql(host string) VersionChecker {
	return mysql{
		host: host,
	}
}
