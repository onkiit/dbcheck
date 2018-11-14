package check

import (
	"database/sql"
	"fmt"
)

type mysql struct {
	DB   *sql.DB
	host string
}

type activeClient struct {
	id       int     `json:"Id"`
	user     string  `json:"User"`
	host     string  `json:"Host"`
	db       string  `json:"db"`
	command  string  `json:"Command"`
	time     int     `json:"Time"`
	state    string  `json:"State"`
	info     string  `json:"Info"`
	progress float32 `json:"Progress"`
}

func (m *mysql) Version() error {
	rows, err := m.DB.Query("SHOW VARIABLES LIKE '%version%'")
	if err != nil {
		fmt.Println("query ", err)
		return nil
	}

	var version, variable, value string
	version = "MySql "
	for rows.Next() {
		err := rows.Scan(&variable, &value)
		if err != nil {
			fmt.Println("scan", err)
			return err
		}
		version += value + " "
	}

	fmt.Println(version)

	return nil
}

func (m *mysql) ActiveClient() error {
	rows, err := m.DB.Query("SHOW PROCESSLIST")
	if err != nil {
		return err
	}

	var info activeClient
	for rows.Next() {
		_ = rows.Scan(info.id, info.user, info.host, info.db, info.command, info.time, info.state, info.info, info.progress)
	}

	fmt.Println(info)
	return nil
}

func (m *mysql) Dial() error {
	db, err := sql.Open("mysql", m.host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	m.DB = db
	return nil
}

func (m *mysql) GetInfo() error {
	if err := m.Dial(); err != nil {
		fmt.Println(err)
		return err
	}
	if err := m.Version(); err != nil {
		fmt.Println(err)
		return err
	}
	if err := m.ActiveClient(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewMysql(host string) DBChecker {
	return &mysql{
		host: host,
	}
}
