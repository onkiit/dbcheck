package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	DB   *sql.DB
	host string
}

type activeClient struct {
	user    string `json:"USER"`
	host    string `json:"HOST"`
	db      string `json:"DB"`
	command string `json:"COMMAND"`
	state   string `json:"STATE"`
	info    string `json:"INFO"`
	time    string `json:"TIME_MS"`
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
	rows, err := m.DB.Query("SELECT USER, HOST, DB, COMMAND, STATE, INFO, TIME_MS FROM INFORMATION_SCHEMA.PROCESSLIST")
	if err != nil {
		return err
	}

	activeClients := []activeClient{}
	for rows.Next() {
		var c activeClient
		_ = rows.Scan(&c.user, &c.host, &c.db, &c.command, &c.state, &c.info, &c.time)

		activeClients = append(activeClients, c)
	}

	info := fmt.Sprintf("active_client(s): %d \n", len(activeClients))
	for _, v := range activeClients {
		if v.state == "" {
			v.state = "-"
		}
		if v.time == "" {
			v.time = "0"
		}
		if v.info == "" {
			v.info = "-"
		}
		info += " ____________\n User: " + v.user + "\n Host: " + v.host + "\n Command: " + v.command + "\n State: " + v.state + "\n DB Name: " + v.db + "\n Running Queries: " + v.info + "\n Time: " + v.time + "\n"
	}
	fmt.Println(info)
	return nil
}

func (m *mysql) Health() error {
	fmt.Println("health_status: \n")

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
	if err := m.Health(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// func NewMysql(host string) DBChecker {
// 	return &mysql{
// 		host: host,
// 	}
// }
