package database

import "database/sql"

type Psql struct {
	Host string
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

func NewPsql(host string) VersionCheck {
	return &Psql{
		Host: host,
	}
}
